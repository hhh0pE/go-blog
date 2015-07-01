package actions
import (
    "net/http"
    "github.com/hhh0pE/go-blog/models"
    "fmt"
)

func Api(w http.ResponseWriter, r *http.Request) (*models.Page, int) {
    fmt.Println("Api hander!")

    return &models.Page{}, 200
}