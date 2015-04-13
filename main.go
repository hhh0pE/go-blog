// blog.lesnoy.name project main.go
package main

import (
	"./models"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
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
	next_post_link := posts.GetNextLink(post.Id)
	prev_post_link := posts.GetPrevLink(post.Id)

	if !is_post_exist {
		http.NotFound(w, r)
		return
	}

	render("post", struct {
		Post      models.Post
		Prev_link string
		Next_link string
	}{post, prev_post_link, next_post_link}, w)
}

func main() {
	fmt.Println("starting server..")
	r := mux.NewRouter()

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))
	r.HandleFunc("/", rootAction)
	r.HandleFunc("/post/{post_url}/", viewAction)

	err := http.ListenAndServe(":9000", r)
	if err != nil {
		fmt.Println("Error serving port 9000")
		fmt.Println(err)
	}
}
