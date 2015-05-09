package models

import "github.com/hhh0pE/go-blog/models/db"

type Post struct {
    db.Page
}

func (p Post) Next() Post {
    next_post := Post{}
    db.Connection.Table("pages").Where("id > ? and parent_id = ?", p.ID, p.Parent_id).First(&next_post)


    return next_post
}

func (p Post) Prev() Post {
    prev_post := Post{}

    db.Connection.Table("pages").Where("id < ? and parent_id = ?", p.ID, p.Parent_id).Order("id DESC").First(&prev_post)


    return prev_post
}

func GetAllPosts() *[]Post {
    all_posts := []Post{}

    db.Connection.Table("pages").Where("template_id = 2").Order("updated_at desc").Order("id").Find(&all_posts)

    return &all_posts
}