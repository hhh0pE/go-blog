package routing

import (
	"bytes"
	"github.com/hhh0pE/go-blog/models/db"
	"html/template"
	"net/http"
)

func render(model interface{}, temp *db.Template, writer http.ResponseWriter) {

	tmpl, err := template.ParseFiles(temp.ToStrings()...)
	if err != nil {
		panic("Error when parsing templates " + temp.ToString() + "`. Error message: " + err.Error())
	}

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "layout", model)

	if err != nil {
		panic("Error when execute template" + err.Error())
	}

	writer.Write(buf.Bytes())
}
