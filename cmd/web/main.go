package main

import (
	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/web"
)

func main() {
	web.Serve(config.Get().RecipeDirectory, config.Get().Web.Port)
}
