package piglatin

import (
	"strings"
)

// Vowels and special cases
var vowels = "aeiou"
var specialCases = []string{"xr", "yt"}

// Sentence translates a full sentence into Pig Latin.
func Sentence(sentence string) string {
	words := strings.Fields(sentence)
	for i, word := range words {
		words[i] = translateWord(word)
	}
	return strings.Join(words, " ")
}

// translateWord translates a single word into Pig Latin.
func translateWord(word string) string {
	// Rule 1:  Word begins with a vowel or "xr" or "yt"
	if startsWithVowelOrSpecialCase(word) {
		return word + "ay"
	}

	firstVowelIndex := -1
	for i, char := range word {
		if strings.ContainsRune(vowels, char) {
			firstVowelIndex = i
			break
		}
	}
	if firstVowelIndex == -1 {
		// Rule 4: Words starting with consonant(s) followed by 'y'
		if index := strings.Index(word, "y"); index > 0 {
			return word[index:] + word[:index] + "ay"
		}

		return word + "ay" // No vowels in the word
	}

	// Rule 3: Words with 'qu' after some consonants
	if word[firstVowelIndex] == 'u' && word[firstVowelIndex-1] == 'q' {
		return word[firstVowelIndex-1+2:] + word[:firstVowelIndex-1+2] + "ay"
	}

	// Rule 4: Words starting with consonant(s) followed by 'y'
	if word[firstVowelIndex-1] == 'y' {
		return word[firstVowelIndex:] + word[:firstVowelIndex] + "ay"
	}

	// Rule 2: Words starting with consonant(s)
	return word[firstVowelIndex:] + word[:firstVowelIndex] + "ay"
}

// Helper function to check if a word starts with a vowel or a special case.
func startsWithVowelOrSpecialCase(word string) bool {
	lowerWord := strings.ToLower(word)
	if strings.ContainsRune(vowels, rune(lowerWord[0])) {
		return true
	}
	for _, special := range specialCases {
		if strings.HasPrefix(lowerWord, special) {
			return true
		}
	}
	return false
}
