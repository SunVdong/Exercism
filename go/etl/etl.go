package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int)
	for point, arr := range in {
		for _, v := range arr {
			res[strings.ToLower(v)] = point
		}
	}

	return res
}
