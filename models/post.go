package models

import (
	"time"
    "fmt"
)

type Post struct {
	Id           int
	Content      string
	Title        string
	Url          string
	Date         time.Time
	Time         time.Time
	Viewed_count int
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

func (p Post) GetByUrl(url string) (Post, bool) {
    var result Post
    query := DB.Where("url = ?", url).First(&result)

    return result, query.RowsAffected > 0
}

func (p Post) GetNextLink(id int) string {
    var link string
    row := DB.Table("posts").Select("url").Where("id > ?", id).Limit(1).Row()
    err := row.Scan(&link)

    if err != nil {
        fmt.Printf("Next link of post ID#%d not found\n", id)
    }

    return link
}

func (p Post) GetPrevLink(id int) string {
    var link string
    row := DB.Table("posts").Select("url").Where("id < ?", id).Limit(1).Row()
    err := row.Scan(&link)

    if err != nil {
        fmt.Printf("Prev link of post ID#%d not found\n", id)
    }

    return link
}