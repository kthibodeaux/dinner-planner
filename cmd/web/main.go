package main

import (
	"log"

	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/web"
)

func main() {
	config := config.LoadConfig()

	log.Println("Recipes:", config.RecipeDirectory)

	web.Serve(config.RecipeDirectory, config.Web.Port)
}
