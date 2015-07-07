package routing

import (
	"bytes"
	"fmt"
	"github.com/hhh0pE/go-blog/models"
	"html/template"
	"net/http"
)

var func_map template.FuncMap

func init() {
	func_map = make(template.FuncMap)
	func_map["ViewedCountText"] = func(count int) string {
		if count == 0 {
			return ""
		}
		switch {
		case count > 1 && count < 5:
			return fmt.Sprintf("%d %s", count, "просмотра")
		case count > 5 && count < 20:
			return fmt.Sprintf("%d %s", count, "просмотров")
		case count%10 == 1:
			return fmt.Sprintf("%d %s", count, "просмотр")
		case count%10 > 1 && count%10 < 5:
			return fmt.Sprintf("%d %s", count, "просмотра")
		}
		return fmt.Sprintf("%d %s", count, "просмотров")
	}
	func_map["OtherPostsInThisCategory"] = models.OtherPostsInThisCategory
	func_map["AllPostsInCategory"] = models.AllPostsInCategory
	func_map["AllCategories"] = models.AllCategories
	func_map["AllPosts"] = models.AllPosts
    func_map["IsProduction"] = models.IsProduction
    func_map["CurrentUser"] = func() *models.User {
        return CurrentUser
    }

}

func render(model interface{}, temp *models.Template, writer http.ResponseWriter) {

	tmpl := template.New(temp.Name).Funcs(func_map)

	var err error
	tmpl, err = tmpl.ParseFiles(temp.ToStrings()...)
	if err != nil {
		panic("Error when parsing templates " + temp.ToString() + "`. Error message: " + err.Error())
	}

	tmpl.Funcs(func_map)

	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, "layout", model)

	if err != nil {
		panic("Error when execute template" + err.Error())
	}

	writer.Write(buf.Bytes())
}
