package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) (res []string) {
	subject = strings.ToLower(subject)
	l := len(subject)

	for _, item := range candidates {
		item_lower := strings.ToLower(item)
		if item_lower == subject {
			continue
		}
		if l != len(item_lower) {
			continue
		}
		for _, letter := range subject {
			item_lower = strings.Replace(item_lower, string(letter), "", 1)
		}
		if item_lower == "" {
			res = append(res, item)
		}
	}

	return
}
