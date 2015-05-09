// blog.lesnoy.name project main.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hhh0pE/go-blog/models"
	"net/http"
	//"github.com/justinas/alice" // lightweight middleware
	//	"net/url"
	"html/template"
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

	render(post_page, "post", w)
}

func categoryAction(w http.ResponseWriter, r *http.Request) {

}

var router = mux.NewRouter()

func main() {
	fmt.Println("starting server..")

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))
	router.HandleFunc("/", rootAction).Name("index")
	router.HandleFunc("/{category}/", categoryAction).Name("category")
	router.HandleFunc("/{category}/{post_url}/", postAction).Name("post")

	err := http.ListenAndServe(":9000", router)
	if err != nil {
		fmt.Println("Error serving port 9000")
		fmt.Println(err)
	}
}
