package actions

import (
	"github.com/hhh0pE/go-blog/models"
	"net/http"
	"strconv"
    "time"
)

func Login(w http.ResponseWriter, r *http.Request) (*models.Page, int) {

    login_page := models.Page{}

	r.ParseForm()
	if len(r.PostForm) == 2 { // login and password
		username := r.PostForm["login"][0]
		password := r.PostForm["password"][0]

		if user, exist := models.UserAuthorize(username, password); exist {
			http.SetCookie(w, &http.Cookie{Name: "UserID", Value: strconv.Itoa(user.ID), Path: "/"})
            login_page.ParentID = 3
            return &login_page, 301
		}
	}

	login_page.Template = &models.Template{File: "admin/login.html", Parent: &models.Template{File: "admin/layout.html"}}

	return &login_page, 200
}

func Logout(w http.ResponseWriter, r *http.Request) (*models.Page, int) {

    http.SetCookie(w, &http.Cookie{Name: "UserID", Value: "", Path: "/", Expires:time.Now().AddDate(-1,-1,-1)})

    login_page := models.Page{ParentID:3}
    return &login_page, 301
}
