package models

import "github.com/hhh0pE/go-blog/models/db"

type Page interface {
	GetTemplates() []string
	Permalink() string
	Parent() *db.Page
}
