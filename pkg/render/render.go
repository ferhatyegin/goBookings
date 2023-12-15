package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate is where program parses HTML templates and it's content to be written by the response writer
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}
