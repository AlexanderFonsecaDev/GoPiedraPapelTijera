package templates

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	TemplateDirectory = "templates/"
	TemplateBase      = TemplateDirectory + "base.html"
)

func LoadTemplate(w http.ResponseWriter, nameTemplate string, data any) {
	templateToLoad := template.Must(template.ParseFiles(TemplateBase, TemplateDirectory+nameTemplate+".html"))
	err := templateToLoad.ExecuteTemplate(w, "base", data)
	if err != nil {
		detailError := fmt.Sprintf("Error rendering template %s , error : %s", nameTemplate, err.Error())
		http.Error(w, detailError, http.StatusInternalServerError)
		return
	}
}
