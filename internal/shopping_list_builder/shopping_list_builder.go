package shoppingListBuilder

import (
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

type ShoppingListBuilder struct {
	SelectedIndex       int
	ShoppingListRecipes []*ShoppingListBuilderRecipe
}

func NewShoppingList(allRecipes []*recipe.Recipe, recipes []*recipe.Recipe) []*ShoppingListBuilderRecipe {
	shoppingListRecipes := make([]*ShoppingListBuilderRecipe, 0)
	for _, recipe := range recipes {
		shoppingListRecipes = append(shoppingListRecipes, addRecipeAndDependants(allRecipes, recipe, 0)...)
	}
	return shoppingListRecipes
}

func addRecipeAndDependants(allRecipes []*recipe.Recipe, recipe *recipe.Recipe, depth int) []*ShoppingListBuilderRecipe {
	shoppingListRecipes := make([]*ShoppingListBuilderRecipe, 0)
	shoppingListRecipe := ShoppingListBuilderRecipe{
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

func (s *ShoppingListBuilder) HandleDown() {
	if s.SelectedIndex >= len(s.ShoppingListRecipes)-1 {
		return
	}
	s.SelectedIndex++
}

func (s *ShoppingListBuilder) HandleUp() {
	if s.SelectedIndex <= 0 {
		return
	}
	s.SelectedIndex--
}

func (s *ShoppingListBuilder) Toggle() {
	s.ShoppingListRecipes[s.SelectedIndex].Include = !s.ShoppingListRecipes[s.SelectedIndex].Include
}
