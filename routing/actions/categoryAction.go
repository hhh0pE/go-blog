package actions

import (
	"github.com/gorilla/mux"
	"github.com/hhh0pE/go-blog/models"
	"github.com/hhh0pE/go-blog/models/db"
	"net/http"
)

func Category(w http.ResponseWriter, r *http.Request) (models.Page, int) {

	vars := mux.Vars(r)

	category := models.Category{}
	var err error
	category.Page, err = db.GetPageByUrl(vars["category"])
	if err != nil {
		return nil, 404
	}

	category.Posts = category.GetAllPosts()

	return category, 200
}
