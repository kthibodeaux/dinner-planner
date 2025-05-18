package shoppingListBuilder

import (
	"fmt"
	"strings"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

const indent = 2

type ShoppingListBuilderRecipe struct {
	Include bool
	Depth   int
	Recipe  *recipe.Recipe
}

func (s *ShoppingListBuilderRecipe) String() string {
	selectionMarker := " "
	if s.Include {
		selectionMarker = "x"
	}

	prefix := fmt.Sprintf("%s[%s] ", strings.Repeat(" ", s.Depth*indent), selectionMarker)

	return prefix + s.Recipe.Name
}

func (s *ShoppingListBuilderRecipe) dependentRecipes(allRecipes []*recipe.Recipe) []*ShoppingListBuilderRecipe {
	dependencies := make([]*ShoppingListBuilderRecipe, 0)

	for _, part := range s.Recipe.Parts {
		for _, ingredient := range part.Ingredients {
			if ingredient.RecipeID != "" {
				for _, recipe := range allRecipes {
					if recipe.ID == ingredient.RecipeID {
						shoppingListRecipe := &ShoppingListBuilderRecipe{
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
