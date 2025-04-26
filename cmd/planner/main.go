package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kthibodeaux/dinner-planner/internal/config"
	"github.com/kthibodeaux/dinner-planner/internal/recipe"
)

func main() {
	config := config.LoadConfig()

	recipes := recipe.Load(config.RecipeDirectory)
	fmt.Printf("Loaded %d recipes", len(recipes))
	fmt.Println()
	dates := loadDates(config.StartDate)

	fmt.Println("Dates:", dates)
}

func loadDates(startDate string) []time.Time {
	date, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Fatal(err)
	}

	dates := make([]time.Time, 7)
	for i := range 7 {
		dates[i] = date.AddDate(0, 0, i)
	}

	return dates
}
