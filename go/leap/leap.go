// Package gives leap year finding powers
package leap

// IsLeapYear determines if the given year is a leap year or not
func IsLeapYear(y int) bool {
	return y%4 == 0 && (y%100 != 0 || y%400 == 0)
}
