package zebra

import (
	"strconv"
)

type Relation int

type Element struct {
	Position   int
	Properties []string
}

type Puzzle struct {
	Position    []Element
	Properties  map[string]*Element
	Candidates  [][]string
	Constraints []Constraint
}

type Constraint struct {
	P string
	Q string
	R Relation
}

const (
	RelationSame Relation = iota
	RelationNext
	RelationImmediateLeft
)

func (p *Puzzle) checkConstraint() bool {
	for _, c := range p.Constraints {
		if p.Properties[c.P] == nil || p.Properties[c.Q] == nil {
			continue
		}
		switch c.R {
		case RelationSame:
			if p.Properties[c.P] != p.Properties[c.Q] {
				return false
			}
		case RelationNext:
			if v := p.Properties[c.P].Position - p.Properties[c.Q].Position; v != 1 && v != -1 {
				return false
			}
		case RelationImmediateLeft:
			if p.Properties[c.P].Position-p.Properties[c.Q].Position != -1 {
				return false
			}
		}
	}
	return true
}

func (p *Puzzle) fillProperty(level, i int) bool {
	if level == len(p.Candidates) {
		return p.checkConstraint()
	}
	if i == len(p.Position) {
		if !p.checkConstraint() {
			return false
		}
		return p.fillProperty(level+1, 0)
	}
	for j := i; j < len(p.Position); j++ {
		p.Candidates[level][i], p.Candidates[level][j] = p.Candidates[level][j], p.Candidates[level][i]
		p.Position[i].Properties[level] = p.Candidates[level][i]
		p.Properties[p.Candidates[level][i]] = &p.Position[i]
		r := p.fillProperty(level, i+1)
		p.Candidates[level][i], p.Candidates[level][j] = p.Candidates[level][j], p.Candidates[level][i]
		if r {
			return true
		}
		p.Position[j].Properties[level] = ""
		delete(p.Properties, p.Candidates[level][j])
	}
	return false
}

func Solve(candidates [][]string, constraints []Constraint) *Puzzle {
	p, q := len(candidates), len(candidates[0])
	r := &Puzzle{
		Position:    make([]Element, q),
		Properties:  make(map[string]*Element),
		Candidates:  candidates,
		Constraints: constraints,
	}
	for i := range r.Position {
		r.Position[i].Position = i
		r.Position[i].Properties = make([]string, p)
		r.Properties[strconv.Itoa(i)] = &r.Position[i]
	}
	if r.fillProperty(0, 0) {
		return r
	}
	return nil
}

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

func SolvePuzzle() Solution {
	var r Solution

	candidates := [][]string{
		{"Red", "Green", "Ivory", "Yellow", "Blue"},
		{"Englishman", "Spaniard", "Ukrainian", "Norwegian", "Japanese"},
		{"Dog", "Snails", "Fox", "Horse", "Zebra"},
		{"Coffee", "Tea", "Milk", "Orange Juice", "Water"},
		{"Old Gold", "Kools", "Chesterfields", "Lucky Strike", "Parliaments"},
	}

	constraints := []Constraint{
		{"Englishman", "Red", RelationSame},
		{"Spaniard", "Dog", RelationSame},
		{"Coffee", "Green", RelationSame},
		{"Ukrainian", "Tea", RelationSame},
		{"Ivory", "Green", RelationImmediateLeft},
		{"Old Gold", "Snails", RelationSame},
		{"Kools", "Yellow", RelationSame},
		{"Milk", "2", RelationSame},
		{"Norwegian", "0", RelationSame},
		{"Chesterfields", "Fox", RelationNext},
		{"Kools", "Horse", RelationNext},
		{"Lucky Strike", "Orange Juice", RelationSame},
		{"Japanese", "Parliaments", RelationSame},
		{"Norwegian", "Blue", RelationNext},
	}

	p := Solve(candidates, constraints)
	if p != nil {
		r.DrinksWater = p.Properties["Water"].Properties[1]
		r.OwnsZebra = p.Properties["Zebra"].Properties[1]
	}
	return r
}
