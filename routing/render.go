package routing

import (

    "net/http"
    "fmt"
    "html/template"
)


func render(model interface{}, templates []string, writer http.ResponseWriter) {

    fmt.Println("rendering..")
    template, err := template.ParseFiles(templates...)
    if err != nil {
        panic("Error when parsing templates " + templates[0] + "`. Error message: " + err.Error())
    }

    // place for middlewares

    err = template.ExecuteTemplate(writer, "layout", model)
    if err != nil {
        panic("Error when execute template" + err.Error())
    }
}