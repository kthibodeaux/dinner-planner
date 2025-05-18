package shoppingList

import (
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

type ShoppingList struct {
	SelectedIndex       int
	ShoppingListRecipes []*ShoppingListRecipe
}

func NewShoppingList(allRecipes []*recipe.Recipe, recipes []*recipe.Recipe) []*ShoppingListRecipe {
	shoppingListRecipes := make([]*ShoppingListRecipe, 0)
	for _, recipe := range recipes {
		shoppingListRecipes = append(shoppingListRecipes, addRecipeAndDependants(allRecipes, recipe, 0)...)
	}
	return shoppingListRecipes
}

func addRecipeAndDependants(allRecipes []*recipe.Recipe, recipe *recipe.Recipe, depth int) []*ShoppingListRecipe {
	shoppingListRecipes := make([]*ShoppingListRecipe, 0)
	shoppingListRecipe := ShoppingListRecipe{
		Include: true,
		Depth:   depth,
		Recipe:  recipe,
	}

	shoppingListRecipes = append(shoppingListRecipes, &shoppingListRecipe)
	dependencies := shoppingListRecipe.dependentRecipes(allRecipes)
	if len(dependencies) > 0 {
		for _, dependency := range dependencies {
			shoppingListRecipes = append(shoppingListRecipes, addRecipeAndDependants(allRecipes, dependency.Recipe, depth+1)...)
		}
	}

	return shoppingListRecipes
}

func (s *ShoppingList) HandleDown() {
	if s.SelectedIndex >= len(s.ShoppingListRecipes)-1 {
		return
	}
	s.SelectedIndex++
}

func (s *ShoppingList) HandleUp() {
	if s.SelectedIndex <= 0 {
		return
	}
	s.SelectedIndex--
}

func (s *ShoppingList) Toggle() {
	s.ShoppingListRecipes[s.SelectedIndex].Include = !s.ShoppingListRecipes[s.SelectedIndex].Include
}
