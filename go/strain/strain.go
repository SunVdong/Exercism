package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) (res Ints) {
	for _, val := range i {
		if filter(val) {
			res = append(res, val)
		}
	}

	return
}

func (i Ints) Discard(filter func(int) bool) (res Ints) {
	for _, val := range i {
		if !filter(val) {
			res = append(res, val)
		}
	}

	return res
}

func (l Lists) Keep(filter func([]int) bool) (res Lists) {
	for _, val := range l {
		if filter(val) {
			res = append(res, val)
		}
	}

	return res
}

func (s Strings) Keep(filter func(string) bool) (res Strings) {
	for _, val := range s {
		if filter(val) {
			res = append(res, val)
		}
	}

	return res
}
