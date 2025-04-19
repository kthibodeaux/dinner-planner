package main

import (
	"testing"
	"time"
)

func TestSunday(t *testing.T) {
	layout := "2006-01-02"

	currentDate, _ := time.Parse(layout, "2025-04-15")
	result := sunday(currentDate)
	expected := "2025-04-13"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	currentDate, _ = time.Parse(layout, "2025-04-13")
	result = sunday(currentDate)
	expected = "2025-04-13"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	currentDate, _ = time.Parse(layout, "2025-04-12")
	result = sunday(currentDate)
	expected = "2025-04-06"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestLoadDates(t *testing.T) {
	startDate := "2025-04-13"
	expectedDates := []string{
		"2025-04-13",
		"2025-04-14",
		"2025-04-15",
		"2025-04-16",
		"2025-04-17",
		"2025-04-18",
		"2025-04-19",
	}

	dates := loadDates(startDate)
	for i, date := range dates {
		if date.Format("2006-01-02") != expectedDates[i] {
			t.Errorf("Expected %s, got %s", expectedDates[i], date.Format("2006-01-02"))
		}
	}
}
