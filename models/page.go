package models

import (
	"html/template"
	"net/url"
	"regexp"
	"strings"
	"time"
    "fmt"
)

type Page struct {
	ID, ParentID, TemplateID                   int
	Url, Title, Description, Keywords, Content string
	Created_at, Updated_at                     time.Time
	ViewedCount                                int
	Image                                      string
	Template                                   *Template
}

//
//type IPage interface {
//	GetTemplate() *Template
//	Permalink() string
//	Parent() *Page
//}

func (p Page) PageType() string {
	var page_type string
	Connection.Table("templates").Select("name").Where("id = ?", p.TemplateID).Row().Scan(&page_type)

	return page_type
}

func (p Page) Parent() *Page {
	if p.ParentID == 0 {
		return nil
	}

	parent_page := Page{}
	query := Connection.Where("id = ?", p.ParentID).First(&parent_page)

	if query.RowsAffected == 0 {
		return nil
	}

	return &parent_page
}

func (p Page) Children() *[]Page {
	children := []Page{}
	query := Connection.Where("parent_id = ?", p.ID).Find(&children)

	if query.RowsAffected == 0 {
		return nil
	}

	return &children
}

func (p Page) Permalink() string {
	if p.Url == "/" {
		return p.Url
	}
	if p.ParentID > 0 {
		return p.Parent().Url + "/" + url.QueryEscape(p.Url)
	}

	return url.QueryEscape(p.Url)
}

func (p Page) HTMLContent() template.HTML {
	return template.HTML(p.Content)
}

func (p Page) MetaDescription() string {
	return strings.Replace(p.Description, "\n", " ", -1)
}

func (p Page) HTMLDescription() template.HTML {
	r, _ := regexp.Compile("(?mi)(.*?)$")
	html_description := r.ReplaceAllString(p.Description, "<p>$1</p>")
	//    fmt.Println(html_description)
	return template.HTML(html_description)
}

func (p Page) GetTemplate() *Template {
	if p.Template != nil {
		return p.Template
	}

	temp := Template{}
	Connection.Table("templates").Where("id = ?", p.TemplateID).First(&temp)
	p.Template = &temp

	if temp.ParentID > 0 {
		temp2 := Template{}
		Connection.Table("templates").Where("id = ?", temp.ParentID).First(&temp2)
		p.Template.Parent = &temp2
	}

	return p.Template
}

func (p Page) CodeBG() template.HTML {
	regxp, _ := regexp.Compile("(?ims)<pre.*?>.*?</pre>")
	code_blocks := strings.Join(regxp.FindAllString(p.Content, -1), "\n<br />\n")
	return template.HTML(code_blocks)
}

func (p *Page) AfterUpdate() (err error) {
    fmt.Println("after update!")
    fmt.Println(p)
	p.Parent().Updated_at = p.Updated_at
    p.Parent().Save()
	return nil
}

func (p *Page) Save() {
	Connection.Save(&p)
}
