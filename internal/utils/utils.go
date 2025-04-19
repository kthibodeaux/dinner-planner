package utils

import (
	"regexp"
	"strings"
)

var slugRegex = regexp.MustCompile(`[^a-z0-9]+`)

func Slugify(s string) string {
	s = strings.ToLower(s)
	s = slugRegex.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}
