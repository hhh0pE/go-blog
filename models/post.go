package models

import (
	"fmt"
	"html/template"
	"time"
	"unicode/utf8"
)

type Post struct {
    Page
	Title, Content string
	CreationDate   time.Time
    Date time.Time
    Time time.Time
	Viewed_count   int
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

func (p Post) Permalink() string {
	return p.Url
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
	DB.Order("date desc").Find(&results)

	return results
}

func (p Post) GetByUrl(url string) error {
    p.Page.GetByUrl(url)
	DB.Where("url = ?", url).First(&p)
	//    fmt.Println(result.Title)

	return nil
}

func (p Post) GetNextPost() Post {
	var result Post
	query := DB.Select("url, title").Where("id > ?", p.ID).Limit(1).First(&result)

	fmt.Println(p.ID)
	fmt.Println(query.Error)
	return result
}

func (p Post) GetPrevPost() Post {
	var result Post
	DB.Select("url, title").Where("id < ?", p.ID).Last(&result)

	return result
}

//
//func (p Post) AfterFind() (err error) {
//    var router = mux.CurrentRoute()
//    fmt.Println(router.Get("post"))
//    fmt.Printf("After find: %+v", p)
//    return nil
//}
