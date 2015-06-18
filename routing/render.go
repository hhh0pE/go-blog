package routing

import (
	"bytes"
	"html/template"
	"net/http"
    "github.com/hhh0pE/go-blog/models/db"
    "fmt"
)

func render(model interface{}, temp *db.Template, writer http.ResponseWriter) {

	template, err := template.ParseFiles(temp.ToStrings()...)
	if err != nil {
		panic("Error when parsing templates " + temp.ToString() + "`. Error message: " + err.Error())
	}

	var buf bytes.Buffer
	err = template.ExecuteTemplate(&buf, "layout", model)

	if err != nil {
		panic("Error when execute template" + err.Error())
	}

	writer.Write(buf.Bytes())
}
