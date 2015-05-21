package models

import "github.com/hhh0pE/go-blog/models/db"

type Category struct {
	db.Page
	Posts []Post
}

func (c Category) GetAllPosts() []Post {
	posts := []Post{}
	db.Connection.Table("pages").Where("parent_id = ?", c.ID).Order("updated_at desc").Find(&posts)

	return posts
}

func GetAllCategories() []Category {
	categories := []Category{}
	db.Connection.Table("pages").Where("template_id = 3").Find(&categories)

	return categories
}

func (c Category) PageType() string {
	return "category"
}

func (c Category) AfterUpdate() (err error) {
	c.Parent().Updated_at = c.Updated_at
	return nil
}
