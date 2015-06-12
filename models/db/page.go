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
	ID, Parent_id, Template_id                 int
	Url, Title, Description, Keywords, Content string
	Created_at, Updated_at                     time.Time
	ViewedCount                                int
	Image                                      string
}

func (p *Page) GetByUrl(url_to_find string) error {
	query := Connection.Where("url = ?", url_to_find).First(&p)

	if err := query.Error; err != nil {
		return errors.New("Error getting page with url " + url_to_find + ". Not found in DB")
	}

	return nil
}

func (p *Page) GetByID(id_to_find int) error {
	query := Connection.First(&p, id_to_find)

	if err := query.Error; err != nil {
		return errors.New("Error getting page with id#" + fmt.Sprintf("%d", id_to_find) + ". Not found in DB")
	}

	return nil
}

func (p Page) Parent() *Page {
	if p.Parent_id == 0 {
		return nil
	}

	parent_page := Page{}
	query := Connection.Where("id = ?", p.Parent_id).First(&parent_page)

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
	if p.Parent_id > 0 {
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

func(p Page) GetTemplates() []string {
    temp := Template{}
    tstrings := []string{}
    Connection.Table("templates").Where("id = ?", p.Template_id).First(&temp)

    tstrings = append(tstrings, "templates/"+temp.File)
    for temp.ParentID > 0 {
        Connection.Table("templates").Where("id = ?", temp.ParentID).First(&temp)
        tstrings = append(tstrings, "templates/"+temp.File)
    }

    return tstrings
}

func (p Page) AfterUpdate() (err error) {
	p.Parent().Updated_at = p.Updated_at
	return nil
}

func (p Page) Save() {
	Connection.Save(p)
}
