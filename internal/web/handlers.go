package web

import (
	"net/http"
	"sort"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
	"github.com/kthibodeaux/dinner-planner/internal/utils"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "index.tmpl", categories)
}

func categoryHandler(w http.ResponseWriter, r *http.Request) {
	categoryRecipes := make([]*recipe.Recipe, 0)
	for _, recipe := range recipes {
		if utils.Slugify(recipe.CookbookCategory) == r.PathValue("id") {
			categoryRecipes = append(categoryRecipes, recipe)
		}
	}

	sort.Slice(categoryRecipes, func(i, j int) bool {
		return categoryRecipes[i].Name < categoryRecipes[j].Name
	})

	data := struct {
		Category string
		Recipes  []*recipe.Recipe
	}{
		Category: categories[r.PathValue("id")],
		Recipes:  categoryRecipes,
	}

	render(w, "category.tmpl", data)
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "recipe.tmpl", recipes[r.PathValue("id")])
}
