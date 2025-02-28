package binarysearch

func SearchInts(list []int, key int) int {
	start := 0
	end := len(list) - 1

	for start <= end {
		curr := (start + end) / 2
		if key == list[curr] {
			return curr
		} else if key > list[curr] {
			start = curr + 1
		} else {
			end = curr - 1
		}
	}

	return -1
}
