// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	by4, by100, by400 := year%4, year%100, year%400

	if by4 == 0 {
		if by400 == 0 {
			return true
		}

		if by100 == 0 {
			return false
		}

		return true
	}

	return false
}
