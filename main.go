// blog.lesnoy.name project main.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Page struct {
	title         string
	template      []byte
	template_path string
}

var counter int = 0

func rootAction(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "<html><head><title>Golang site!</title></head><body>Hello Golang site! <br /> It's %d hits alredy!</body></html>", counter)

	fmt.Println(r)
}

func main() {
	fmt.Println("starting server..")
	r := mux.NewRouter()

	//r.Host("blog.lesnoy.name")

	r.HandleFunc("/", rootAction)

	err := http.ListenAndServe(":9000", r)
	if err != nil {
		fmt.Println("Error serving port 9000")
		fmt.Println(err)
	}
}
