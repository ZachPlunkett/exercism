// Package acronym ...
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate ...
func Abbreviate(s string) string {
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "-", " ")

	re := regexp.MustCompile("[^a-zA-Z-]")
	split := strings.Split(s, " ")

	var acronym string
	for _, s := range split {
		s = re.ReplaceAllString(s, "")

		if len(s) > 0 {
			acronym += string(s[0])
		}
	}

	return acronym
}
