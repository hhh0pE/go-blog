// blog.lesnoy.name project main.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
    "html/template"
//    "./models"
)

func render(name string, data interface{}, w http.ResponseWriter) {
    t, err := template.ParseFiles("templates/"+name+".tmpl")
    if err != nil {
        panic("Error when parsing template `"+name+"`. Error message: "+err.Error())
    }

    t.Execute(w, data)

}

func rootAction(w http.ResponseWriter, r *http.Request) {

    render("index", struct{Text, Name string}{"index page!", "Vlad"}, w)
}

func blogAction(w http.ResponseWriter, r *http.Request) {
    render("post", struct{Text, Name string}{"posts page!", "Vlad"}, w)
}

func main() {
	fmt.Println("starting server..")
	r := mux.NewRouter()

	r.HandleFunc("/", rootAction)
    r.HandleFunc("/blog/", blogAction)

	err := http.ListenAndServe(":9000", r)
	if err != nil {
		fmt.Println("Error serving port 9000")
		fmt.Println(err)
	}
}
