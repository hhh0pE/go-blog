package routing

import (
	"bytes"
	"html/template"
	"net/http"
)

func render(model interface{}, templates []string, writer http.ResponseWriter) {

	template, err := template.ParseFiles(templates...)
	if err != nil {
		panic("Error when parsing templates " + templates[0] + "`. Error message: " + err.Error())
	}

	var buf bytes.Buffer
	err = template.ExecuteTemplate(&buf, "layout", model)

	if err != nil {
		panic("Error when execute template" + err.Error())
	}

	writer.Write(buf.Bytes())
}
