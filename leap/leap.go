/*Package leap verifies if given a year, report if it is a leap year in the Gregorian calendar*/
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	// is leap
	// - IF evenly divided by 4 but NOT leap if evenly divided by 100
	// - IF evenly divided  by 400
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
