package stringset

import (
	"fmt"
	"sort"
	"strings"
)

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set struct {
	data []string
}

func New() Set {
	return Set{data: nil}
}

func NewFromSlice(l []string) Set {
	s := Set{data: []string{}}
	for _, e := range l {
		if !s.Has(e) {
			s.data = append(s.data, e)
		}
	}
	return s
}

func (s Set) String() string {
	if len(s.data) != 0 {
		return fmt.Sprintf("{\"" + strings.Join(s.data, "\", \"") + "\"}")
	}
	return "{}"
}

func (s Set) IsEmpty() bool {
	return len(s.data) == 0
}

func (s Set) Has(elem string) bool {
	for _, v := range s.data {
		if v == elem {
			return true
		}
	}

	return false
}

func (s *Set) Add(elem string) {
	if !s.Has(elem) {
		s.data = append(s.data, elem)
	}
}

func Subset(s1, s2 Set) bool {
	for _, e1 := range s1.data {
		if !s2.Has(e1) {
			return false
		}
	}

	return true
}

func Disjoint(s1, s2 Set) bool {
	for _, e1 := range s1.data {
		if s2.Has(e1) {
			return false
		}
	}

	return true
}

func Equal(s1, s2 Set) bool {
	sort.Strings(s1.data)
	sort.Strings(s2.data)
	return strings.Join(s1.data, ",") == strings.Join(s2.data, ",")
}

func Intersection(s1, s2 Set) Set {
	s := Set{data: []string{}}
	for _, e := range s1.data {
		if s2.Has(e) {
			s.data = append(s.data, e)
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := Set{data: []string{}}
	for _, e := range s1.data {
		if !s2.Has(e) {
			s.data = append(s.data, e)
		}
	}

	return s
}

func Union(s1, s2 Set) Set {
	s := Set{data: s1.data}
	for _, e := range s2.data {
		s.Add(e)
	}
	return s
}
