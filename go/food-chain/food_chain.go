package foodchain

import (
	"fmt"
	"strings"
)

// Data for the song
var animals = []string{
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

var reactions = []string{
	"",
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
	"She's dead, of course!",
}

func Verse(v int) string {
	if v == 8 {
		return swalloed(animals[v-1]) + reactions[v-1]
	}

	var sb strings.Builder
	sb.WriteString(swalloed(animals[v-1]))
	if v > 1 {
		sb.WriteString(reactions[v-1] + "\n")
	}
	for i := v - 1; i > 0; i-- {
		sb.WriteString(swallowedToCatch(animals[i], animals[i-1]))
	}
	sb.WriteString("I don't know why she swallowed the fly. Perhaps she'll die.")
	return sb.String()
}

func Verses(start, end int) string {
	var sb strings.Builder
	for i := start; i < end; i++ {
		sb.WriteString(Verse(i) + "\n\n")
	}
	sb.WriteString(Verse(end))
	return sb.String()
}

func Song() string {
	return Verses(1, 8)
}

func swalloed(what string) string {
	return fmt.Sprintf("I know an old lady who swallowed a %s.\n", what)
}

func swallowedToCatch(swallowed, toCatch string) string {
	if toCatch == "spider" {
		return fmt.Sprintf("She swallowed the %s to catch the %s that wriggled and jiggled and tickled inside her.\n", swallowed, toCatch)
	}
	return fmt.Sprintf("She swallowed the %s to catch the %s.\n", swallowed, toCatch)
}
