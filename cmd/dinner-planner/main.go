package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kthibodeaux/dinner-planner/internal/recipe"
	"github.com/sanity-io/litter"
)

type config struct {
	directory string
	startDate string
}

func main() {
	var config config
	flag.StringVar(&config.directory, "directory", defaultPath(), "Directory containing recipes")
	flag.StringVar(&config.startDate, "startdate", sunday(time.Now()), "Start date in yyyy-mm-dd format (defaults to current week's Sunday)")
	flag.Parse()

	recipes := recipe.Load(config.directory)
	fmt.Printf("Loaded %d recipes", len(recipes))
	for _, recipe := range recipes {
		litter.Dump(recipe)
		fmt.Println()
		fmt.Println()
	}
	loadDates(config.startDate)
}

func defaultPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(home, "recipes")
}

func sunday(currentDate time.Time) string {
	offset := int(currentDate.Weekday())
	date := currentDate.AddDate(0, 0, -offset)

	return date.Format("2006-01-02")
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
