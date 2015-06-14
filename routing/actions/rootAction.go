package actions

import (
    "net/http"
    "github.com/hhh0pE/go-blog/models"
    "time"
    "fmt"
)

func Root(r *http.Request) (models.HomePage, *http.Request) {



    homepage := models.HomePage{}
    homepage.GetByUrl("/")

    homepage.Posts = models.GetAllPosts()
    homepage.Categories = models.GetAllCategories()

    cook, err := r.Cookie("last-viewed")
    if err == nil && cook != nil {
        lv_date, lv_err := time.Parse(time.RFC3339, cook.Value)
        if lv_err != nil {
            fmt.Println(lv_err.Error())
        }
        homepage.LastViewedDate = lv_date
    } else {
        homepage.LastViewedDate = time.Now() // by default
    }

    r.AddCookie(&http.Cookie{Name: "last-viewed", Value: time.Now().Format(time.RFC3339), Expires: time.Now().AddDate(1, 0, 0)})

    return homepage, r
}