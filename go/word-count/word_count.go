package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

// var r = regexp.MustCompile("([0-9]+)|([a-z]+'?[a-z]+)|([a-z]+)")
var r = regexp.MustCompile("[a-z0-9]+(['][a-z0-9]+)?")

func WordCount(phrase string) Frequency {
	m := make(Frequency, 0)
	phrase = strings.ToLower(phrase)
	arr := r.FindAllString(phrase, -1)
	for _, item := range arr {
		m[item] += 1
	}

	return m
}
