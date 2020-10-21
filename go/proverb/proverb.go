// Package proverb has a function to make a riddle
package proverb

func bodyString(w1, w2 string) string {
	return "For want of a " + w1 + " the " + w2 + " was lost."
}

func lastString(w1 string) string {
	return "And all for the want of a " + w1 + "."
}

// Proverb produces a rhyme given a slice of strings
func Proverb(rhyme []string) []string {
	rhymeLen := len(rhyme)

	// Edge cases for 0 and 1 word
	if rhymeLen == 0 {
		return []string{}
	}

	var sentences []string
	for i := 0; i < len(rhyme)-1; i++ {
		sentences = append(sentences, bodyString(rhyme[i], rhyme[i+1]))
	}
	return append(sentences, lastString(rhyme[0]))
}
