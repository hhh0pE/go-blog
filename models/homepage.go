package models

import (
	"github.com/hhh0pE/go-blog/models/db"
	"time"
)

type HomePage struct {
	db.Page
	Posts          []Post
	Categories     []Category
	LastViewedDate time.Time
}

func (hp HomePage) PageType() string {
	return "homepage"
}
