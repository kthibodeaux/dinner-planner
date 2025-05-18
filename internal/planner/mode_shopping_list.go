package planner

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
	shoppingList "github.com/kthibodeaux/dinner-planner/internal/shopping_list"
	shoppingListBuilder "github.com/kthibodeaux/dinner-planner/internal/shopping_list_builder"
)

func (dp *dinnerPlan) viewModeShoppingList() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		dp.shoppingListSelectColumn(),
		dp.shoppingListShowColumn(),
	)
}

func (dp *dinnerPlan) shoppingListSelectColumn() string {
	title := "Shopping List"
	header := lipgloss.NewStyle().Render(title + "\n")

	content := make([]string, 0)
	content = append(content, header)
	for i, slr := range dp.shoppingList.ShoppingListRecipes {
		if i == dp.shoppingList.SelectedIndex {
			content = append(content, styleSelected.Render(slr.String()))
		} else {
			content = append(content, slr.String())
		}
	}

	return dp.stylePaneBorder(borderForce).
		Width(dp.size.width/2 - 2).
		Height(dp.size.height - 2).
		Render(strings.Join(content, "\n"))
}

func (dp *dinnerPlan) shoppingListShowColumn() string {
	recipes := make([]*recipe.Recipe, 0)
	for _, slr := range dp.shoppingList.ShoppingListRecipes {
		if slr.Include {
			recipes = append(recipes, slr.Recipe)
		}
	}

	list := shoppingList.NewShoppingList(recipes)

	return dp.stylePaneBorder(borderForceHidden).
		Width(dp.size.width/2 - 2).
		Height(dp.size.height - 2).
		Render(strings.Join(list, "\n"))
}

func (dp *dinnerPlan) prepareShoppingList() {
	shoppingListRecipes := make([]*shoppingListBuilder.ShoppingListBuilderRecipe, 0)
	for i := 1; i < len(dp.recipeLists); i++ {
		for _, r := range dp.recipeLists[i].Recipes {
			slr := shoppingListBuilder.NewShoppingList(dp.recipeLists[0].Recipes, []*recipe.Recipe{r})
			shoppingListRecipes = append(shoppingListRecipes, slr...)
		}
	}

	dp.shoppingList = &shoppingListBuilder.ShoppingListBuilder{
		ShoppingListRecipes: shoppingListRecipes,
	}
}
