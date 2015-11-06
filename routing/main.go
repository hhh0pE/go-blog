package routing

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hhh0pE/go-blog/models"
)

var router *mux.Router

func init() {
	router = mux.NewRouter()
}

var CurrentUser *models.User

func RouteFunc(pattern string, action func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(pattern, action)
}

func Route(pattern string, action func(http.ResponseWriter, *http.Request) (*models.Page, int)) {
	router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {

		CurrentUser = nil
		if cook, err := r.Cookie("UserID"); err == nil {
			uid, _ := strconv.Atoi(cook.Value)
			CurrentUser, _ = models.GetUserByID(uid)
		}

		pageModel, code := action(w, r)

		if code == 301 {
			http.Redirect(w, r, "/"+pageModel.Parent().Permalink()+"/", 301)
			return
		}

		if code == 404 {
			if CurrentUser == nil || CurrentUser.Role != "admin" {
				http.NotFound(w, r)
				return
			}
		}

		if pageModel != nil {
			render(pageModel, pageModel.GetTemplate(), w)
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
