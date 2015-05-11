package models
import "github.com/hhh0pE/go-blog/models/db"

type Category struct {
    db.Page
    Posts *[]Post
}

func (c Category) GetAllPosts() *[]Post {
    posts := []Post{}
    db.Connection.Table("pages").Where("parent_id = ?", c.ID).Find(&posts)

    return &posts
}

func GetAllCategories() *[]Category {
    categories := []Category{}
    db.Connection.Table("pages").Where("template_id = 3").Find(&categories)

    return &categories
}