package models

import (
	"fmt"
	"html/template"
	"time"
	"unicode/utf8"
)

type Post struct {
    Page
    Page_id int
    ID  int
	Title, Content string
	CreationDate   time.Time
    Date time.Time
    Time time.Time
	Viewed_count   int
}

func (p *Post) AfterFind() (err error) {
    p.Page.ID = p.Page_id
    return
}

func (p Post) HTMLContent() template.HTML {

	return template.HTML(p.Content)
}

func (p Post) Description() template.HTML {
	str := p.Content
	var shortDescription []rune
	for i := 0; i < 500; i++ {
		symbol, width := utf8.DecodeRuneInString(str)
		shortDescription = append(shortDescription, symbol)

		str = str[width:]
	}

	return template.HTML(shortDescription)
}

func (p Post) GetLast() Post {
	var result Post
	DB.First(&result)

	return result
}

func (p Post) GetFirst() Post {
	var result Post
	DB.Last(&result)

	return result
}

func (p Post) GetAll() []Post {
	var results []Post
    pages, _ := DB.Raw("SELECT * FROM pages WHERE id IN(SELECT page_id FROM posts)").Rows()
    for pages.Next() {
        new_page := Post{}
        pages.Scan(&new_page.Page.ID, &new_page.Parent_id, &new_page.Metatitle, &new_page.Metadescription, &new_page.Metakeywords, &new_page.Page.Url)

        results = append(results, new_page)
    }
	return results
}

func (p *Post) GetByUrl(url string) error {
    p.Page.GetByUrl(url)
    DB.Where("page_id = ?", p.Page.ID).First(&p)

	return nil
}

func (p Post) GetNextPost() Post {
	var result Post
	query := DB.Select("id, page_id, title").Where("id > ? AND page_id<>0", p.ID).Limit(1).First(&result)

	fmt.Println(p.ID)
	fmt.Println(query.Error)
	return result
}

func (p Post) GetPrevPost() Post {
	var result Post
	DB.Select("id, page_id, title").Where("id < ? AND page_id<>0", p.ID).Last(&result)

	return result
}

//
//func (p Post) AfterFind() (err error) {
//    var router = mux.CurrentRoute()
//    fmt.Println(router.Get("post"))
//    fmt.Printf("After find: %+v", p)
//    return nil
//}
