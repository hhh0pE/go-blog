package actions

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/feeds"
	"github.com/hhh0pE/go-blog/models"
	"time"
)

func Rss(w http.ResponseWriter, r *http.Request) {
	categoryName := mux.Vars(r)["category"]

	// TODO: rss without category - for all site

	var page *models.Page
	var children *[]models.Page
	page, exist := models.GetPageByUrl(categoryName)

	if !exist {
		http.Error(w, "Category not found", 404)
		return
	}

	children = page.ChildrenByCreatedAt()
	if children == nil {
		http.Error(w, "Category doesn't contains any post", 404)
		return
	}

	url := "http://lesnoy.name/"

	feed := &feeds.Feed{
		Title: page.Title,
		Link: &feeds.Link{Href: url+page.Permalink()+"/"},
		Description: page.Description,
		Author: &feeds.Author{"Vladislav Lesnoy", "vlad@lesnoy.name"},
		Created: time.Now(),
		Copyright: "Copyrights by Vladislav Lesnoy",
	}

	for _, page := range *children {
		item := feeds.Item{
			Title: page.Title,
			Link: &feeds.Link{Href: url+page.Permalink()+"/"},
			Description: page.Description,
			Created: page.Created_at,
		}
		fmt.Printf("%#v\n", item)
		fmt.Println(item)
		feed.Add(&item)
	}

	feed.WriteRss(w)
}