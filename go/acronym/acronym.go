// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
    "strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
    replaced := strings.ReplaceAll(s, "-", " ")

	// 分割成单词
	words := strings.Fields(replaced)

	// 提取每个单词的第一个字母并转为大写
	var result []rune
	for _, word := range words {
		if len(word) > 0 {
			firstRune := rune(word[0])
			// 转为大写
			uppercaseFirstRune := unicode.ToUpper(firstRune)
			result = append(result, uppercaseFirstRune)
		}
	}

	// 将结果转换为字符串
	output := string(result)
	return output
}
