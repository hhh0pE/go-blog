package actions

import (
    "github.com/joeguo/sitemap"
    "github.com/hhh0pE/go-blog/models"
    "fmt"
)

func BuildSitemap() {
    categories := models.GetAllCategories()
    posts := models.GetAllPosts()

    homepage := models.HomePage{}
    homepage.GetByUrl("/")

    sitemap_items := make([]*sitemap.Item, len(categories)+len(posts)+1)

    sitemap_items[0] = &sitemap.Item{"http://lesnoy.name" + homepage.Permalink(), homepage.Updated_at, "daily", 0.8}

    for i, c := range categories {
        sitemap_items[i+1] = &sitemap.Item{Loc: "http://lesnoy.name/" + c.Permalink() + "/", LastMod: c.Updated_at, Changefreq: "daily", Priority: 1}
    }

    for i, p := range posts {
        sitemap_items[i+1+len(categories)] = &sitemap.Item{Loc: "http://lesnoy.name/" + p.Permalink() + "/", LastMod: p.Updated_at, Changefreq: "weekly", Priority: 0.5}
    }

    err := sitemap.SiteMap("public/sitemap.xml", sitemap_items)
    if err != nil {
        fmt.Println(err.Error())
    }
}