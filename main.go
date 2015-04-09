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
	t, err := template.ParseFiles("templates/" + name + ".tmpl")
	if err != nil {
		panic("Error when parsing template `" + name + "`. Error message: " + err.Error())
	}

	t.Execute(w, data)

}

func rootAction(w http.ResponseWriter, r *http.Request) {
    post_model := models.Post{}
    posts := post_model.GetAll()
	render("index", posts, w)
}

func viewAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	post_url := vars["post_url"]

	posts := models.Post{}
	post, exist := posts.GetByUrl(post_url)

    if !exist {
        http.NotFound(w, r)
        return
    }

	render("post", post, w)
}

func main() {
	fmt.Println("starting server..")
	r := mux.NewRouter()

	r.HandleFunc("/", rootAction)
	r.HandleFunc("/{post_url}/", viewAction)

	err := http.ListenAndServe(":9000", r)
	if err != nil {
		fmt.Println("Error serving port 9000")
		fmt.Println(err)
	}
}
