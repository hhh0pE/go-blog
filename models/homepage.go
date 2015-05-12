package models

import "github.com/hhh0pE/go-blog/models/db"

type HomePage struct {
	db.Page
	Posts []Post
    Categories []Category
}

func (hp HomePage) PageType() string {
    return "homepage"
}