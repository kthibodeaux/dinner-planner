package main

import (
	"log"

	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/web"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Recipes:", config.RecipeDirectory)

	web.Serve(config.RecipeDirectory, config.Web.Port)
}
