// blog.lesnoy.name project main.go
package main

import (
	"fmt"

	"github.com/hhh0pE/go-blog/routing"
	"github.com/hhh0pE/go-blog/routing/actions"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:9002", nil))
	}()
	host := "localhost:9001"
	fmt.Println("starting server " + host + "..")

	routing.RouteDirectory("assets")

	routing.Route("/", actions.Root)
	routing.Route("/login/", actions.Login)
	routing.Route("/logout/", actions.Logout)
	routing.RouteFunc("/rss/{category}/", actions.Rss)
	routing.Route("/api/{entity}/", actions.Api)
	routing.Route("/{category}/", actions.Category)
	routing.Route("/{category}/{post_url}/", actions.Post)


//	routing.Route("/rss/", actions.Rss)

	routing.RouteFile("/sitemap.xml", "public/sitemap.xml")
	routing.RouteFile("/robots.txt", "public/robots.txt")

	err := http.ListenAndServe(host, routing.Router())
	if err != nil {
		fmt.Println("Error serving " + host)
		fmt.Println(err)
	}
}
