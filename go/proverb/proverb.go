// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	len := len(rhyme)
	if len == 0 {
		return []string{}
	}

	var proverb []string = make([]string, len)
	for i := 0; i < len-1; i++ {
		proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
	}
	proverb[len-1] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])

	return proverb
}
