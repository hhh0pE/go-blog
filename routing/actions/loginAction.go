package actions

import (
    "net/http"
    "github.com/hhh0pE/go-blog/models"
    "github.com/hhh0pE/go-blog/models/db"
//    "github.com/gorilla/mux"
)

func Login(w http.ResponseWriter, r *http.Request) (models.Page, int) {

    login_page := db.Page{}
    login_page.Template = &db.Template{File:"admin/login.html", Parent:&db.Template{File:"admin/layout.html"}}
//    vars := mux.Vars(r)

    return login_page, 200
}