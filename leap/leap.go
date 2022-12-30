// Package leap provides leap year operetations
package leap

// IsLeapYear returns true if the provided year is a leap year, false otherwise.
func IsLeapYear(input int) bool {
	if input%400 == 0 {
		return true
	}
	if input%100 == 0 {
		return false
	}
	return input%4 == 0
}
