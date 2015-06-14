package routing

import (
    "bytes"
    "net/http"
    "html/template"
    "compress/gzip"
    "fmt"
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

    var zbuf bytes.Buffer
    zip := gzip.NewWriter(&zbuf)
    zip.Write(buf.Bytes())
    zip.Close()
    
    writer.Header().Add("Content-Encoding", "gzip")
    writer.Header().Add("Content-Type", "text/html")
    writer.Header().Add("Content-Length", fmt.Sprintf("%d", zbuf.Len()))
    writer.Write(zbuf.Bytes())
//    writer.Write(buf.Bytes())
}