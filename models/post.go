package models

import (
    "time"
    "fmt"
)

type Post struct {
    id int
    content string
    title string
    url string
    date time.Time
    time time.Time
    viewed_count int
}

func (p Post) GetLast() Post {
    result := Post{}
    err := DB.QueryRow("SELECT * FROM posts ORDER BY date,time DESC LIMIT 1").Scan(&result.id, &result.content, &result.viewed_count, &result.title, &result.url, &result.date, &result.time)
    if err != nil {
        fmt.Println("Error when selecting GetLast() from table `blog`. "+err.Error())
    }

    return result
}

func (p Post) GetFirst() Post {
    result := Post{}
    err := DB.QueryRow("SELECT * FROM posts ORDER BY date,time ASC LIMIT 1").Scan(&result.id, &result.content, &result.viewed_count, &result.title, &result.url, &result.date, &result.time)
    if err != nil {
        fmt.Println("Error when selecting GetFirst() from table `blog`. "+err.Error())
    }

    return result
}

func (p Post) GetAll() []Post {
    rows, err := DB.Query("SELECT * FROM posts ORDER BY date,time DESC")
    if err != nil {
        fmt.Println("Error when selecting GetAll() from table `blog`. "+err.Error())
    }

    var results []Post

    for rows.Next() {
        result := Post{}
        rows.Scan(&result.id, &result.content, &result.viewed_count, &result.title, &result.url, &result.date, &result.time)
        results = append(results, result)
    }

    return results
}