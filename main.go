// blog.lesnoy.name project main.go
package main

import (
	"fmt"

	"github.com/hhh0pE/go-blog/routing"
	"github.com/hhh0pE/go-blog/routing/actions"
	"net/http"
)

func main() {
	fmt.Println("starting server..")

	routing.RouteDirectory("assets")

	routing.Route("/", actions.Root)
    routing.Route("/login/", actions.Login)
	routing.Route("/{category}/", actions.Category)
	routing.Route("/{category}/{post_url}/", actions.Post)

	routing.RouteFile("/sitemap.xml", "public/sitemap.xml")
	routing.RouteFile("/robots.txt", "public/robots.txt")

	err := http.ListenAndServe(":9001", routing.Router())
	if err != nil {
		fmt.Println("Error serving port 9001")
		fmt.Println(err)
	}
}
