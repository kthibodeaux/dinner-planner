package plan

import (
	"time"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

type Plan struct {
	Date    time.Time
	Recipes []*recipe.Recipe
}
