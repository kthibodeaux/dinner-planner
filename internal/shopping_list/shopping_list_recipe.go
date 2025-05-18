package shoppingList

import (
	"fmt"
	"strings"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

const indent = 2

type ShoppingListRecipe struct {
	Include bool
	Depth   int
	Recipe  *recipe.Recipe
}

func (s *ShoppingListRecipe) String() string {
	selectionMarker := " "
	if s.Include {
		selectionMarker = "x"
	}

	prefix := fmt.Sprintf("%s[%s] ", strings.Repeat(" ", s.Depth*indent), selectionMarker)

	return prefix + s.Recipe.Name
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
