package actions

import (
	"github.com/gorilla/mux"
	"github.com/hhh0pE/go-blog/models"
	"net/http"
)

func Category(w http.ResponseWriter, r *http.Request) (*models.Page, int) {

	vars := mux.Vars(r)

	page, exist := models.GetPageByUrl(vars["category"])

	if !exist {
		return nil, 404
	}

	return page, 200
}
