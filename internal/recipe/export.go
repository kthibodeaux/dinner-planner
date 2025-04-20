package recipe

import (
	"html/template"
	"os"
)

func (recipe *Recipe) ExportToCookbookHTML(filename string) {
	template, err := template.ParseFiles(
		"templates/cookbook/layout.tmpl",
		"templates/cookbook/recipe.tmpl",
	)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	if err := template.Execute(outFile, recipe); err != nil {
		panic(err)
	}
}
