package recipe

import (
	"html/template"
	"os"
)

func (recipe *Recipe) ExportToCookbookHTML(filename string) error {
	template, err := template.ParseFiles(
		"templates/cookbook/layout.tmpl",
		"templates/cookbook/recipe.tmpl",
	)
	if err != nil {
		return err
	}

	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := template.Execute(outFile, recipe); err != nil {
		return err
	}

	return nil
}
