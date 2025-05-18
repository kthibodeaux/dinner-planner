package shoppingList

import (
	"slices"
	"strings"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

type quantities []string
type ingredients map[string]quantities

func NewShoppingList(recipes []*recipe.Recipe) []string {
	list := make(ingredients, 0)

	for _, recipe := range recipes {
		for _, part := range recipe.Parts {
			for _, ingredient := range part.Ingredients {
				key := strings.ToLower(ingredient.Name)
				if _, exists := list[key]; !exists {
					list[key] = make(quantities, 0)
				}

				quantity := ingredient.Quantity
				if quantity != "" {
					if ingredient.Unit != "" {
						quantity += " " + ingredient.Unit
					}
					list[key] = append(list[key], quantity)
				}
			}
		}
	}

	content := make([]string, 0)
	for name, quantities := range list {
		line := name
		if len(quantities) > 0 {
			line = line + " (" + strings.Join(quantities, ", ") + ")"
		}
		content = append(content, line)
	}

	slices.Sort(content)

	return content
}
