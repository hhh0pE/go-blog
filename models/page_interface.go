package models

import "github.com/hhh0pE/go-blog/models/db"

type Page interface {
	GetTemplate() *db.Template
	Permalink() string
	Parent() *db.Page
	SetUser(*db.User) db.Page
}
