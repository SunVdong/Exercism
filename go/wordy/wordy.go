package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	replacer := strings.NewReplacer(
		"What is ", "",
		" by", "",
		"?", "",
	)
	question = replacer.Replace(question)

	var sum int
	var opt string

	has_left := false
	for _, word := range strings.Fields(question) {
		if n, err := strconv.Atoi(word); err == nil {
			if !has_left {
				sum = n
				has_left = true
				continue
			}
			if has_left {
				switch opt {
				case "plus":
					sum += n
				case "minus":
					sum -= n
				case "divided":
					sum /= n
				case "multiplied":
					sum *= n
				default:
					return 0, false
				}
				opt = ""
			}
		} else {
			if !has_left {
				return 0, false
			}
			if opt != "" {
				return 0, false
			}
			opt = word
		}
	}

	if opt != "" {
		return 0, false
	}

	return sum, true
}
