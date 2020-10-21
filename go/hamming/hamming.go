// Package hamming determines hamming distance of DNA strands
package hamming

import (
	"fmt"
	"strings"
)

const (
	lenErr = "DNA strands need to be of equal length: received %s-%d, %s-%d"
)

// Distance finds the hamming distance of two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf(lenErr, a, len(a), b, len(b))
	}

	a = strings.ToLower(a)
	b = strings.ToLower(b)

	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}
