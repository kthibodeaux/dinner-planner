package main

import (
	"fmt"

	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
	"github.com/kthibodeaux/dinner-planner/internal/utils"
)

func main() {
	config := config.LoadConfig()

	recipes := recipe.Load(config.RecipeDirectory)
	dates := utils.DatesForWeekStartingOn(config.StartDate)

	fmt.Println("Recipes:", len(recipes))
	fmt.Println("Dates:", dates)
}
