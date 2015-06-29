package actions

import (
	"github.com/hhh0pE/go-blog/models"
	"net/http"
	//    "github.com/gorilla/mux"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) (*models.Page, int) {

	r.ParseForm()

	if len(r.PostForm) == 2 { // login and password
		username := r.PostForm["login"][0]
		password := r.PostForm["password"][0]

		if user, exist := models.UserAuthorize(username, password); exist {
			http.SetCookie(w, &http.Cookie{Name: "UserID", Value: strconv.Itoa(user.ID), Path: "/"})
		}

	}
	login_page := models.Page{}
	login_page.Template = &models.Template{File: "admin/login.html", Parent: &models.Template{File: "admin/layout.html"}}
	//    vars := mux.Vars(r)

	return &login_page, 200
}
