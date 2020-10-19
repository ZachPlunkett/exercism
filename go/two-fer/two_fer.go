// Package twofer returns a string for you and me.
package twofer

// ShareWith returns a string like "One for X, one for me."
// X is your input name. If it is empty "you" will be used.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
