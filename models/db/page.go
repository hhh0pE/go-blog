package db

import (
	"errors"
	"fmt"
	"html/template"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type Page struct {
	Id, ParentID, TemplateID                 int
	Url, Title, Description, Keywords, Content string
	Created_at, Updated_at                     time.Time
	ViewedCount                                int
	Image                                      string
    Template                                    *Template
}

func GetPageByUrl(url_to_find string) (Page, error) {
    p := Page{}
    query := Connection.Where("url = ?", url_to_find).First(&p)

    if err := query.Error; err != nil {
        return p, errors.New("Error getting page with url " + url_to_find + ". Not found in DB")
    }

    return p, nil
}

func GetPageByID(id_to_find int) (Page, error) {
    p := Page{}
    query := Connection.First(&p, id_to_find)

    if err := query.Error; err != nil {
        return p, errors.New("Error getting page with id#" + fmt.Sprintf("%d", id_to_find) + ". Not found in DB")
    }

    return p, nil
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
	query := Connection.Where("parent_id = ?", p.Id).Find(&children)

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

func (p Page) AfterUpdate() (err error) {
	p.Parent().Updated_at = p.Updated_at
	return nil
}

func (p Page) Save() {
	Connection.Save(p)
}
