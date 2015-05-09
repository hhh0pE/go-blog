package db

import (
	"errors"
	"fmt"
	"time"
    "html/template"
)

type Page struct {
	ID, Parent_id, Template_id                 int
	Url, Title, Description, Keywords, Content string
	Created_at, Updated_at                     time.Time
}

func (p *Page) GetByUrl(url_to_find string) error {
    query := Connection.Where("url = ?", url_to_find).First(&p)

    if err := query.Error; err != nil {
        return errors.New("Error getting page with url " + url_to_find+". Not found in DB")
    }

    return nil
}

func (p *Page) GetByID(id_to_find int) error {
    query := Connection.First(&p, id_to_find)

    if err := query.Error; err != nil {
        return errors.New("Error getting page with id#" + fmt.Sprintf("%d", id_to_find)+". Not found in DB")
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
    if p.Parent_id > 0 {
        return p.Parent().Url + "/" + p.Url
    }

    return p.Url
}

func (p Page) HTML() template.HTML {
    return template.HTML(p.Content)
}