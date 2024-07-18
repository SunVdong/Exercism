package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Define the Garden type here.

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

type Garden struct {
	diagram string
	Dm      map[string][]string
}

var plants = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	arr := strings.Split(diagram, "\n")
	if len(arr) < 3 {
		return nil, errors.New("wrong")
	}
	if len(children) <= 0 {
		return nil, errors.New("wrong")
	}

	m := make(map[string][]string)
	sort.Strings(sort.StringSlice(children))

	for i, v := range children {
		m[v] = []string{
			plants[rune(arr[1][i*2])],
			plants[rune(arr[1][i*2+1])],
			plants[rune(arr[2][i*2])],
			plants[rune(arr[2][i*2+1])],
		}
	}
	gd := Garden{diagram: diagram, Dm: m}
	return &gd, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if v, bool := g.Dm[child]; bool {
		return v, true
	}
	return nil, false
}
