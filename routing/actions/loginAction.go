package actions

import (
	"github.com/hhh0pE/go-blog/models"
	"github.com/hhh0pE/go-blog/models/db"
	"net/http"
	//    "github.com/gorilla/mux"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) (models.Page, int) {

	r.ParseForm()

	user := db.User{}
	if len(r.PostForm) == 2 { // login and password
		username := r.PostForm["login"][0]
		password := r.PostForm["password"][0]

		if user.Authorize(username, password) {
			http.SetCookie(w, &http.Cookie{Name: "UserID", Value: strconv.Itoa(user.ID), Path: "/"})
		}

	}
	login_page := db.Page{}
	login_page.Template = &db.Template{File: "admin/login.html", Parent: &db.Template{File: "admin/layout.html"}}
	//    vars := mux.Vars(r)

	return login_page, 200
}
