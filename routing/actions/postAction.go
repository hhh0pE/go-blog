package actions

import (
    "time"
    "net/http"
    "github.com/hhh0pE/go-blog/models"
    "github.com/gorilla/mux"
)

func Post(r *http.Request) (models.Post, *http.Request) {

    vars := mux.Vars(r)
    category_url, post_url := vars["category"], vars["post_url"]

    category_page := models.Category{}
    err := category_page.GetByUrl(category_url)
    if err != nil {
//        http.NotFound(w, r)
//        return
    }

    post_page := models.Post{}
    err = post_page.GetByUrl(post_url)
    if err != nil {
//        http.NotFound(w, r)
//        return
    }

    if post_page.Template_id == 5 { // redirect
        new_post := models.Post{}
        new_post.GetByID(post_page.Parent_id)
//        http.Redirect(w, r, "/"+new_post.Permalink()+"/", 301)
//        return
    }

    cook, cookie_err := r.Cookie("av")
    if cookie_err != nil && cook == nil {
        post_page.ViewedCount++
        post_page.Save()
        r.AddCookie(&http.Cookie{Name: "av", Value: "1", Expires: time.Now().AddDate(1, 0, 0)})
    }

    r.Header.Add("Last-Modified", post_page.Updated_at.Format(time.RFC1123))

    return post_page, r
}