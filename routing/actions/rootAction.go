package actions

import (
    "net/http"
    "github.com/hhh0pE/go-blog/models"
    "time"
    "fmt"
    "github.com/hhh0pE/go-blog/models/db"
)

func Root(w http.ResponseWriter, r *http.Request) (models.Page, int) {

    var err error

    homepage := models.HomePage{}
    homepage.Page, err = db.GetPageByUrl("/")

    if err != nil {
        return homepage, 404
    }

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

    http.SetCookie(w, &http.Cookie{Name: "last-viewed", Value: time.Now().Format(time.RFC3339), Expires: time.Now().AddDate(1, 0, 0)})

    return homepage, 200
}