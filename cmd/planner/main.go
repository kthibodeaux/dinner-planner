package main

import (
	"log"
	"time"

	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/planner"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

func main() {
	recipes, err := recipe.Load()
	if err != nil {
		log.Fatal(err)
	}

	dates, err := datesForWeekStartingOn()
	if err != nil {
		log.Fatal(err)
	}

	planner.Run(recipes, dates)
}

func datesForWeekStartingOn() ([]time.Time, error) {
	date, err := time.Parse("2006-01-02", config.Get().StartDate)
	if err != nil {
		return nil, err
	}

	dates := make([]time.Time, 7)
	for i := range 7 {
		dates[i] = date.AddDate(0, 0, i)
	}

	return dates, nil
}
