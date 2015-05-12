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

func render(model interface{}, template_name string, writer http.ResponseWriter) {
	template, err := template.ParseFiles("templates/"+template_name+".html", "templates/layout.html")
	if err != nil {
		panic("Error when parsing template " + template_name + "`. Error message: " + err.Error())
	}

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

	render(homepage, "index", w)
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

	w.Header().Set("Last-Modified", post_page.Updated_at.Format(time.RFC1123))

	render(post_page, "post", w)
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

	render(category, "category", w)
}

func doSitemap() {
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
}

var router = mux.NewRouter()

func main() {
	fmt.Println("starting server..")

	doSitemap()

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", rootAction).Name("index")
	router.HandleFunc("/{category}/", categoryAction).Name("category")
	router.HandleFunc("/{category}/{post_url}/", postAction).Name("post")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	err := http.ListenAndServe(":9001", router)
	if err != nil {
		fmt.Println("Error serving port 9001")
		fmt.Println(err)
	}
}
