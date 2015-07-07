package actions
import (
    "net/http"
    "github.com/hhh0pE/go-blog/models"
    "fmt"
    "strings"
    "strconv"
)

func Api(w http.ResponseWriter, r *http.Request) (*models.Page, int) {
    err := r.ParseMultipartForm(1000)
    if err != nil {
        fmt.Println("Error when parse POST data: "+err.Error())
    }

    user_cookie, err := r.Cookie("UserID")
    if err != nil {
        fmt.Println("Trying to save data with API from non-user. Error: "+err.Error())
        return nil, 404
    }

    user_id, _ := strconv.Atoi(user_cookie.Value)
    user, _ := models.GetUserByID(user_id)
    if user.Role != "admin" {
        fmt.Println("Trying to save data with API from not admin user. Error: "+err.Error())
        return nil, 404
    }

    url_parts := strings.Split(strings.Trim(r.Referer(), "/"), "/")
    post, _ := models.GetPageByUrl(url_parts[len(url_parts)-1])
    post.Content = r.MultipartForm.Value["content"][0]
    post.Title = r.MultipartForm.Value["title"][0]
    post.Description = r.MultipartForm.Value["description"][0]
    post.Save()
    
    return nil, 200
}