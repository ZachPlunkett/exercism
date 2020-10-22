// Package bob is lazy
package bob

import (
	"strings"
	"unicode"
)

func hasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// Hey ...
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")
	remarkLen := len(remark)
	if remarkLen == 0 {
		return "Fine. Be that way!"
	}
	remarkEnd := remark[remarkLen-1]

	caps := remark == strings.ToUpper(remark)
	letter := hasLetter(remark)

	if remarkEnd == '?' {
		if caps && letter {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else if caps && letter {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
