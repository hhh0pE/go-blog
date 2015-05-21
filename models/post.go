package models

import (
	"fmt"
	"github.com/hhh0pE/go-blog/models/db"
	"html/template"
	"regexp"
	"strings"
)

type Post struct {
	db.Page
}

func (p Post) Next() Post {
	next_post := Post{}
	db.Connection.Table("pages").Where("id > ? and parent_id = ?", p.ID, p.Parent_id).First(&next_post)

	p.Updated_at.Unix()
	return next_post
}

func (p Post) Prev() Post {
	prev_post := Post{}

	db.Connection.Table("pages").Where("id < ? and parent_id = ?", p.ID, p.Parent_id).Order("id DESC").First(&prev_post)

	return prev_post
}

func GetAllPosts() []Post {
	all_posts := []Post{}

	db.Connection.Table("pages").Where("template_id = 2").Order("updated_at desc").Order("id").Find(&all_posts)

	return all_posts
}

func (p Post) GetOtherInCategory() []Post {
	other_in_cat := []Post{}
	db.Connection.Table("pages").Where("parent_id = ? and id <> ?", p.Parent_id, p.ID).Order("viewed_count DESC").Find(&other_in_cat)

	return other_in_cat
}

func (p Post) ViewedCountText() string {
	if p.ViewedCount == 0 {
		return ""
	}
	switch {
	case p.ViewedCount > 1 && p.ViewedCount < 5:
		return fmt.Sprintf("%d %s", p.ViewedCount, "просмотра")
	case p.ViewedCount > 5 && p.ViewedCount < 20:
		return fmt.Sprintf("%d %s", p.ViewedCount, "просмотров")
	case p.ViewedCount%10 == 1:
		return fmt.Sprintf("%d %s", p.ViewedCount, "просмотр")
	case p.ViewedCount%10 > 1 && p.ViewedCount%10 < 5:
		return fmt.Sprintf("%d %s", p.ViewedCount, "просмотра")
	}
	return fmt.Sprintf("%d %s", p.ViewedCount, "просмотров")
}

func (p Post) CodeBG() template.HTML {
	regxp, _ := regexp.Compile("(?ims)<pre.*?>.*?</pre>")

	code_blocks := strings.Join(regxp.FindAllString(p.Content, -1), "\n<br />\n")

	return template.HTML(code_blocks)

}

func (p Post) PageType() string {
	return "post"
}
