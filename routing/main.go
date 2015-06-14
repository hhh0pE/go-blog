package routing

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "github.com/hhh0pE/go-blog/models"
)

var router *mux.Router

func init() {
    router = mux.NewRouter()
}

func Route(pattern string, f func(*http.Request) (models.IPage, *http.Request)) {
    router.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        page_model, req := f(r)
        fmt.Println(req)
        render(page_model, page_model.GetTemplates(), w)


    })
}

func RouteFile(url, path string) {
    router.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Encoding", "gzip")
        http.ServeFile(w, r, path)
    })
}

func RouteDirectory(name string) {
    router.PathPrefix("/"+name+"/").Handler(http.StripPrefix("/"+name, http.FileServer(http.Dir("./"+name+"/"))))
}

func Router() *mux.Router {
    return router
}