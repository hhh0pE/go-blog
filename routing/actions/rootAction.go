package actions

import (
	"fmt"
	"github.com/hhh0pE/go-blog/models"
	"net/http"
	"time"
)

func Root(w http.ResponseWriter, r *http.Request) (*models.Page, int) {

	homepage, homepage_exist := models.GetPageByUrl("/")

	if !homepage_exist {
		return homepage, 404
	}

	cook, err := r.Cookie("last-viewed")
	if err == nil && cook != nil {
		lv_date, lv_err := time.Parse(time.RFC3339, cook.Value)
		if lv_err != nil {
			fmt.Println(lv_err.Error())
		}
		homepage.LastViewedDate = lv_date
	} else {
		homepage.LastViewedDate = time.Now() // by default
	}

	http.SetCookie(w, &http.Cookie{Name: "last-viewed", Value: time.Now().Format(time.RFC3339), Expires: time.Now().AddDate(1, 0, 0)})

	return homepage, 200
}
