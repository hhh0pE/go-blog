package actions
import (
    "net/http"
    "github.com/hhh0pE/go-blog/models"
    "fmt"
    "strings"
)

func Api(w http.ResponseWriter, r *http.Request) (*models.Page, int) {
    err := r.ParseMultipartForm(1000)
    if err != nil {
        fmt.Println("Error when parse POST data: "+err.Error())
    }

    url_parts := strings.Split(strings.Trim(r.Referer(), "/"), "/")
    post, _ := models.GetPageByUrl(url_parts[len(url_parts)-1])
    post.Content = r.MultipartForm.Value["content"][0]
    post.Save()

//    fmt.Println(post.ID)

    return nil, 200
}