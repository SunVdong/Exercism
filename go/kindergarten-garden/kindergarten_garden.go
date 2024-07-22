package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	diagram  string
	children []string
	datamap  map[string][]string
}

var plants = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(diagram) > 0 && diagram[0] != '\n' {
		return nil, errors.New("wrong diagram format")
	}

	if strings.ToUpper(diagram) != diagram {
		return nil, errors.New("invalid cup codes")
	}

	diagram_str := strings.ReplaceAll(diagram, "\n", "")
	if len(children)*4 != len(diagram_str) {
		return nil, errors.New("wrong")
	}

	copyChildren := make([]string, len(children))
	copy(copyChildren, children)
	sort.Strings(copyChildren)

	m := make(map[string][]string, len(children))
	runes := []rune(diagram_str)

	for idx, child := range copyChildren {
		if _, ok := m[child]; ok {
			return nil, errors.New("duplacate name")
		}
		pls := []string{plants[runes[idx*2]], plants[runes[idx*2+1]], plants[runes[(idx+len(children))*2]], plants[runes[(idx+len(children))*2+1]]}
		m[child] = pls
	}

	return &Garden{diagram: diagram_str, children: children, datamap: m}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (g.datamap)[child]
	return plants, ok
}
