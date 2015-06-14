package actions

import (
    "net/http"
    "github.com/hhh0pE/go-blog/models"
    "github.com/gorilla/mux"
)

func Category(r *http.Request) (models.Category, *http.Request) {

    vars := mux.Vars(r)

    category := models.Category{}
    err := category.GetByUrl(vars["category"])
    if err != nil {
//        http.NotFound(w, r)
    }

    category.Posts = category.GetAllPosts()

    return category, r
}