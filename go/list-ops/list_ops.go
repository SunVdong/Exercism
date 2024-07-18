package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, elem := range s {
		initial = fn(initial, elem)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	n := s.Length()
	for n = n - 1; n >= 0; n-- {
		initial = fn(s[n], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	if s.Length() == 0 {
		return s
	}

	var res IntList
	for _, v := range s {
		if fn(v) {
			res = append(res, v)
		}
	}
	return res
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	for i, v := range s {
		s[i] = fn(v)
	}
	return s
}

func (s IntList) Reverse() IntList {
	n := len(s)
	rev := make([]int, n)
	for i := 0; i < n; i++ {
		rev[i] = s[n-1-i]
	}
	return rev
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	// var res IntList
	for _, v := range lists {
		s = append(s, v...)
	}
	return s
}
