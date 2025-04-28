package main

import (
	"log"
	"time"

	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/planner"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	recipes, err := recipe.Load(config.RecipeDirectory)
	if err != nil {
		log.Fatal(err)
	}

	dates, err := datesForWeekStartingOn(config.StartDate)
	if err != nil {
		log.Fatal(err)
	}

	planner.Run(config, recipes, dates)
}

func datesForWeekStartingOn(startDate string) ([]time.Time, error) {
	date, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		return nil, err
	}

	dates := make([]time.Time, 7)
	for i := range 7 {
		dates[i] = date.AddDate(0, 0, i)
	}

	return dates, nil
}
