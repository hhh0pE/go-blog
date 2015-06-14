package actions

import (
    "time"
    "net/http"
    "github.com/hhh0pE/go-blog/models"
    "github.com/gorilla/mux"
    "github.com/hhh0pE/go-blog/models/db"
)

func Post(w http.ResponseWriter, r *http.Request) (models.Page, int) {

    vars := mux.Vars(r)
    category_url, post_url := vars["category"], vars["post_url"]

    var err error

    category_page := models.Category{}
    category_page.Page, err = db.GetPageByUrl(category_url)
    if err != nil {
        return nil, 404
    }

    post_page := models.Post{}
    post_page.Page, err = db.GetPageByUrl(post_url)
    if err != nil {
        return nil, 404
    }

    if post_page.TemplateID == 5 { // redirect
        return post_page, 301
    }

    cook, cookie_err := r.Cookie("av")
    if cookie_err != nil && cook == nil {
        post_page.ViewedCount++
        post_page.Save()
        http.SetCookie(w,&http.Cookie{Name: "av", Value: "1", Expires: time.Now().AddDate(1, 0, 0)} )
    }

    w.Header().Add("Last-Modified", post_page.Updated_at.Format(time.RFC1123))

    return post_page, 200
}