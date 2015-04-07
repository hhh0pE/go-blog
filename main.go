// blog.lesnoy.name project main.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
    "html/template"
    "./models"
)

type Person struct {
    Name string
    Age int
}

type Stat struct {
    Hits int
}

var counter int = 0

const temp = `
Hello, {{.Name}}!
I'm {{.Age}} old.

You just done {{.Hits}} hit!
`

func rootAction(w http.ResponseWriter, r *http.Request) {
    t := template.New("test")
    t, _ = t.Parse(temp)
    t.Execute(w, Person{Name:"Test", Age:22})
	counter++

	fmt.Println(r)
}

func blogAction(w http.ResponseWriter, r *http.Request) {
    posts := models.Post{}
    fmt.Println(posts.GetLast())
    fmt.Println(posts.GetFirst())
    fmt.Println(posts.GetAll())
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
