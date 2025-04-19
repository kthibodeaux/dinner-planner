package web

import (
	"log"
	"net/http"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

var recipes map[string]*recipe.Recipe
var categories map[string]string

func Serve(directory string, port string) {
	recipes = getRecipes(directory)
	categories = getUniqueCategories()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./templates/web/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", indexHandler)
	mux.HandleFunc("GET /categories/{id}", categoryHandler)
	mux.HandleFunc("GET /recipes/{id}", recipeHandler)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
