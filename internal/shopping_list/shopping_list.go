package shoppingList

import (
	"fmt"
	"strings"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

const indent = 2

type ShoppingList struct {
	SelectedIndex       int
	ShoppingListRecipes []*ShoppingListRecipe
}

type ShoppingListRecipe struct {
	Include bool
	Depth   int
	Recipe  *recipe.Recipe
}

func NewShoppingList(allRecipes []*recipe.Recipe, recipes []*recipe.Recipe) []*ShoppingListRecipe {
	shoppingListRecipes := make([]*ShoppingListRecipe, 0)
	for _, recipe := range recipes {
		shoppingListRecipes = append(shoppingListRecipes, addRecipeAndDependants(allRecipes, recipe, 0)...)
	}
	return shoppingListRecipes
}

func (s *ShoppingListRecipe) String() string {
	selectionMarker := " "
	if s.Include {
		selectionMarker = "x"
	}

	prefix := fmt.Sprintf("%s[%s] ", strings.Repeat(" ", s.Depth*indent), selectionMarker)

	return prefix + s.Recipe.Name
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

func (s *ShoppingListRecipe) dependentRecipes(allRecipes []*recipe.Recipe) []*ShoppingListRecipe {
	dependencies := make([]*ShoppingListRecipe, 0)

	for _, part := range s.Recipe.Parts {
		for _, ingredient := range part.Ingredients {
			if ingredient.RecipeID != "" {
				for _, recipe := range allRecipes {
					if recipe.ID == ingredient.RecipeID {
						shoppingListRecipe := &ShoppingListRecipe{
							Include: true,
							Depth:   s.Depth + 1,
							Recipe:  recipe,
						}
						dependencies = append(dependencies, shoppingListRecipe)
					}
				}
			}
		}
	}
	return dependencies
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
