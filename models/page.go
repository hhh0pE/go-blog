package models

import (
	"errors"
    "fmt"
)

type Page struct {
	ID, Parent_id                            int
	Url                                      string
	Metatitle, Metadescription, Metakeywords string
}

func (h *Page) GetByUrl(url string) error {

	query := DB.Table("pages").Where("url = ?", url).First(&h)

	if err := query.Error; err != nil {
		return errors.New("Page with url " + url + " Not Found")
	}

	return nil
}

func (h *Page) GetParent() *Page {
	parent_page := Page{}
	query := DB.Table("pages").Where("id = ?", h.Parent_id).Find(&parent_page)

	if query.RowsAffected == 0 {
		return nil
	}

	return &parent_page
}

func (p Page) Permalink() string {

    fmt.Println("Permalink!")
    if len(p.Url) == 0 {
        DB.Select("parent_id, url").First(&p, p.ID)
    }

    if p.Parent_id > 0 {
        p_page := Page{ID:p.Parent_id}
        return p_page.Permalink()+"/"+p.Url
    }

//    if p.Parent_id != 0 {
//        p_page := Page{}
//
//    }

    return p.Url
}
//
//func (p *Page) GetChildren() []*Page {
//    return &Page{}
//}