package models
import "github.com/hhh0pE/go-blog/models/db"

type Category struct {
    db.Page
}

func GetAllCategories() *[]Category {
    categories := []Category{}
    db.Connection.Table("pages").Where("template_id = 3").Find(&categories)

    return &categories
}