package routing

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hhh0pE/go-blog/models"
	"github.com/hhh0pE/go-blog/models/db"
	"net/http"
	"strconv"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
}

func Route(pattern string, action func(http.ResponseWriter, *http.Request) (models.Page, int)) {
	router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {

		page_model, code := action(w, r)

		cook, err := r.Cookie("UserID")
		if err == nil {
			user := db.User{}
			uid, uid_err := strconv.Atoi(cook.Value)
			fmt.Println(uid)
			if uid_err == nil {
				if user.GetByID(uid) {
					page_model = page_model.SetUser(&user)
				}
				fmt.Printf("User finded! %#v\n", user)
			}
		}

		if code == 404 {
			http.NotFound(w, r)
			return
		}

		if code == 301 {
			http.Redirect(w, r, "/"+page_model.Parent().Permalink()+"/", 301)
			return
		}

		fmt.Printf("%#v", page_model)
		if page_model != nil {
			render(page_model, page_model.GetTemplate(), w)
		}
	})
}

func RouteFile(url, path string) {
	router.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	})
}

func RouteDirectory(name string) {
	router.PathPrefix("/" + name + "/").Handler(http.StripPrefix("/"+name, http.FileServer(http.Dir("./"+name+"/"))))
}

func Router() *mux.Router {
	return router
}
