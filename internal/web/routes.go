package web

import (
	"net/http"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

var recipes map[string]*recipe.Recipe
var categories map[string]string

func Serve(directory string, port string) error {
	var err error
	recipes, err = getRecipes(directory)
	if err != nil {
		return err
	}
	categories = getUniqueCategories()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./templates/web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", indexHandler)
	mux.HandleFunc("GET /categories/{id}", categoryHandler)
	mux.HandleFunc("GET /recipes/{id}", recipeHandler)

	err = http.ListenAndServe(port, mux)
	return err
}
