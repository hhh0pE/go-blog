package models

import (
	"fmt"
	"time"
)

type Post struct {
	id           int
	Content      string
	Title        string
	Url          string
	Date         time.Time
	Time         time.Time
	Viewed_count int
}

func (p Post) GetLast() Post {
	result := Post{}
	err := DB.QueryRow("SELECT * FROM posts ORDER BY date,time DESC LIMIT 1").Scan(&result.id, &result.Content, &result.Viewed_count, &result.Title, &result.Url, &result.Date, &result.Time)
	if err != nil {
		fmt.Println("Error when selecting GetLast() from table `blog`. " + err.Error())
	}

	return result
}

func (p Post) GetFirst() Post {
	result := Post{}
	err := DB.QueryRow("SELECT * FROM posts ORDER BY date,time ASC LIMIT 1").Scan(&result.id, &result.Content, &result.Viewed_count, &result.Title, &result.Url, &result.Date, &result.Time)
	if err != nil {
		fmt.Println("Error when selecting GetFirst() from table `blog`. " + err.Error())
	}

	return result
}

func (p Post) GetAll() []Post {
	rows, err := DB.Query("SELECT * FROM posts ORDER BY date,time DESC")
	if err != nil {
		fmt.Println("Error when selecting GetAll() from table `blog`. " + err.Error())
	}

	var results []Post

	for rows.Next() {
		result := Post{}
		rows.Scan(&result.id, &result.Content, &result.Viewed_count, &result.Title, &result.Url, &result.Date, &result.Time)
		results = append(results, result)
	}

	return results
}

func (p Post) GetByUrl(url string) (Post, bool) {
    sql := "SELECT * FROM posts WHERE url = '"+ url + "'"
	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Println("Error when select post by url. Error: " + err.Error())
	}

    var result Post

    ok := rows.Next()
    if ok {
        rows.Scan(&result.id, &result.Content, &result.Viewed_count, &result.Title, &result.Url, &result.Date, &result.Time)
    }

    return result, ok
}
