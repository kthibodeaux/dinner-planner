package web

import (
	"fmt"
	"html/template"
	"net/http"
)

func render(w http.ResponseWriter, page string, data any) {
	files := []string{
		"templates/web/base.tmpl",
		"templates/web/" + page,
	}

	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Println("Error parsing template:", err.Error())
		http.Error(w, "Template error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Println("Error executing template:", err.Error())
		http.Error(w, "Execution error: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
