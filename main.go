// blog.lesnoy.name project main.go
package main

import (
	"./models"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
)

func render(name string, data interface{}, w http.ResponseWriter) {

	t, err := template.ParseFiles("templates/"+name+".html", "templates/layout.html")
	if err != nil {
		panic("Error when parsing template `" + name + "`. Error message: " + err.Error())
	}

	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		fmt.Println("Error when render page " + name + "; Error: " + err.Error())
	}

}

func rootAction(w http.ResponseWriter, r *http.Request) {
	post_model := models.Post{}
	posts := post_model.GetAll()

	for i := 0; i < len(posts); i++ {
		var err error
		var post_url *url.URL
		post_url, err = router.Get("post").URL("post_url", posts[i].Url)
		if err != nil {
			posts[i].Url = ""
		} else {
			posts[i].Url = post_url.String()
		}

	}

	render("index", struct {
		Title string
		Posts []models.Post
	}{"Lesnoy's blog | main page", posts}, w)
}

func viewAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post_url := vars["post_url"]

	posts := models.Post{}
	post, is_post_exist := posts.GetByUrl(post_url)
	next_post_link, next_err := router.Get("post").URL("post_url", posts.GetNextLink(post.Id))
	prev_post_link, prev_err := router.Get("post").URL("post_url", posts.GetPrevLink(post.Id))

	var next_post_url, prev_post_url string

	if next_err == nil {
		next_post_url = next_post_link.String()
	} else {
		next_post_url = ""
	}

	if prev_err == nil {
		prev_post_url = prev_post_link.String()
	} else {
		prev_post_url = ""
	}

	if !is_post_exist {
		http.NotFound(w, r)
		return
	}

	render("post", struct {
		Post      models.Post
		Prev_link string
		Next_link string
	}{post, prev_post_url, next_post_url}, w)
}

var router = mux.NewRouter()

func main() {
	fmt.Println("starting server..")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", rootAction).Name("index")
	router.HandleFunc("/article/{post_url}/", viewAction).Name("post")

	err := http.ListenAndServe(":9000", router)
	if err != nil {
		fmt.Println("Error serving port 9000")
		fmt.Println(err)
	}
}
