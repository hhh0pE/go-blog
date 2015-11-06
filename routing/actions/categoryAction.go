package actions

import (
	"github.com/gorilla/mux"
	"github.com/hhh0pE/go-blog/models"
	"net/http"
    "time"
)

func Category(w http.ResponseWriter, r *http.Request) (*models.Page, int) {

	vars := mux.Vars(r)

	page, exist := models.GetPageByUrl(vars["category"])

	if !exist {
        return &models.Page{Title:"Новая категория", Created_at: time.Now(), Template:&models.Template{Parent:&models.Template{File:"layout.html"}, Name:"new",File:"category.html"}}, 404
	}

	return page, 200
}
