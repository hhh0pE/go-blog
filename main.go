// blog.lesnoy.name project main.go
package main

import (
	"fmt"

	"github.com/hhh0pE/go-blog/routing"
	"github.com/hhh0pE/go-blog/routing/actions"
	"net/http"
)

func main() {
	host := "localhost:9001"
	fmt.Println("starting server " + host + "..")

	routing.RouteDirectory("assets")

	routing.Route("/", actions.Root)
	routing.Route("/login/", actions.Login)
    routing.Route("/api/{entity}/", actions.Api)
	routing.Route("/{category}/", actions.Category)
	routing.Route("/{category}/{post_url}/", actions.Post)

	routing.RouteFile("/sitemap.xml", "public/sitemap.xml")
	routing.RouteFile("/robots.txt", "public/robots.txt")

	err := http.ListenAndServe(host, routing.Router())
	if err != nil {
		fmt.Println("Error serving " + host)
		fmt.Println(err)
	}
}
