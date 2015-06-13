// blog.lesnoy.name project main.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hhh0pE/go-blog/models"
	"net/http"
	//"github.com/justinas/alice" // lightweight middleware
	//	"net/url"
	"github.com/joeguo/sitemap"
	"html/template"
	"time"
)

func render(model interface{}, templates []string, writer http.ResponseWriter) {

	template, err := template.ParseFiles(templates...)
	if err != nil {
		panic("Error when parsing templates " + templates[0] + "`. Error message: " + err.Error())
	}

	// place for middlewares

	err = template.ExecuteTemplate(writer, "layout", model)
	if err != nil {
		panic("Error when execute template" + err.Error())
	}
}

func rootAction(w http.ResponseWriter, r *http.Request) {

	homepage := models.HomePage{}
	homepage.GetByUrl("/")

	homepage.Posts = models.GetAllPosts()
	homepage.Categories = models.GetAllCategories()

	cook, err := r.Cookie("last-viewed")
	if err == nil && cook != nil {
		lv_date, lv_err := time.Parse(time.RFC3339, cook.Value)
		if lv_err != nil {
			fmt.Println(lv_err.Error())
		}
		homepage.LastViewedDate = lv_date
	} else {
		homepage.LastViewedDate = time.Now() // by default
	}

	http.SetCookie(w, &http.Cookie{Name: "last-viewed", Value: time.Now().Format(time.RFC3339), Expires: time.Now().AddDate(1, 0, 0)})

	render(homepage, []string{"templates/index.html", "templates/layout.html"}, w)
}

func postAction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	category_url, post_url := vars["category"], vars["post_url"]

	category_page := models.Category{}
	err := category_page.GetByUrl(category_url)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	post_page := models.Post{}
	err = post_page.GetByUrl(post_url)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if post_page.Template_id == 5 { // redirect
		new_post := models.Post{}
		new_post.GetByID(post_page.Parent_id)
		http.Redirect(w, r, "/"+new_post.Permalink()+"/", 301)
		return
	}

	cook, cookie_err := r.Cookie("av")
	if cookie_err != nil && cook == nil {
		post_page.ViewedCount++
		post_page.Save()
		http.SetCookie(w, &http.Cookie{Name: "av", Value: "1", Expires: time.Now().AddDate(1, 0, 0)})
	}

	w.Header().Set("Last-Modified", post_page.Updated_at.Format(time.RFC1123))

	render(post_page, post_page.GetTemplates(), w)

	return
}

func categoryAction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	category := models.Category{}
	err := category.GetByUrl(vars["category"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	category.Posts = category.GetAllPosts()

	render(category, category.GetTemplates(), w)
}

func serverSitemap(w http.ResponseWriter, r *http.Request) {
	categories := models.GetAllCategories()
	posts := models.GetAllPosts()

	homepage := models.HomePage{}
	homepage.GetByUrl("/")

	sitemap_items := make([]*sitemap.Item, len(categories)+len(posts)+1)

	sitemap_items[0] = &sitemap.Item{"http://lesnoy.name" + homepage.Permalink(), homepage.Updated_at, "daily", 0.8}

	for i, c := range categories {
		sitemap_items[i+1] = &sitemap.Item{Loc: "http://lesnoy.name/" + c.Permalink() + "/", LastMod: c.Updated_at, Changefreq: "daily", Priority: 1}
	}

	for i, p := range posts {
		sitemap_items[i+1+len(categories)] = &sitemap.Item{Loc: "http://lesnoy.name/" + p.Permalink() + "/", LastMod: p.Updated_at, Changefreq: "weekly", Priority: 0.5}
	}

	err := sitemap.SiteMap("public/sitemap.xml", sitemap_items)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Encoding", "gzip")
	http.ServeFile(w, r, "public/sitemap.xml")
}

var router = mux.NewRouter()

func main() {
	fmt.Println("starting server..")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", rootAction).Name("index")
	router.HandleFunc("/{category}/", categoryAction).Name("category")
	router.HandleFunc("/{category}/{post_url}/", postAction).Name("post")
	//    router.HandleFunc("/admin/", adminAction)

	router.HandleFunc("/sitemap.xml", serverSitemap)
	router.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/robots.txt")
	})

	err := http.ListenAndServe(":9001", router)
	if err != nil {
		fmt.Println("Error serving port 9001")
		fmt.Println(err)
	}
}
