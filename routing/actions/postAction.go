package actions

import (
	"github.com/gorilla/mux"
	"github.com/hhh0pE/go-blog/models"
	"net/http"
	"time"
)

func Post(w http.ResponseWriter, r *http.Request) (*models.Page, int) {

	vars := mux.Vars(r)
	category_url, post_url := vars["category"], vars["post_url"]

	category_page, category_exist := models.GetPageByUrl(category_url)

	if !category_exist {
		return nil, 404
	}

	post_page, post_exist := models.GetPageByUrl(post_url)
	if !post_exist {
        return &models.Page{ParentID:category_page.ID, Title:"Новая запись", Created_at: time.Now(), Template:&models.Template{Parent:&models.Template{File:"layout.html"}, Name:"new",File:"post.html"}}, 404
	}

	if post_page.GetTemplate().ID == 5 { // redirect
		return post_page, 301
	}

	cook, cookie_err := r.Cookie("av")
	if cookie_err != nil && cook == nil {
		post_page.ViewedCount++
		post_page.Save()
		http.SetCookie(w, &http.Cookie{Name: "av", Value: "1", Expires: time.Now().AddDate(1, 0, 0)})
	}

	w.Header().Add("Last-Modified", post_page.Updated_at.Format(time.RFC1123))

	return post_page, 200
}
