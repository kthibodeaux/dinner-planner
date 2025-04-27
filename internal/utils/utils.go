package utils

import (
	"log"
	"regexp"
	"strings"
	"time"
)

var slugRegex = regexp.MustCompile(`[^a-z0-9]+`)

func Slugify(s string) string {
	s = strings.ToLower(s)
	s = slugRegex.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

func DatesForWeekStartingOn(startDate string) []time.Time {
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
