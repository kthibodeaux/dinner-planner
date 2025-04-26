package main

import "strings"

func (dp dinnerPlan) View() string {
	values := make([]string, 0)
	for _, date := range dp.dates {
		values = append(values, date.Format("2006-01-02"))
	}
	return strings.Join(values, "\n")
}
