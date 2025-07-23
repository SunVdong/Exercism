package transpose

func Transpose(input []string) []string {
	maxLen := 0
	for _, s := range input {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

	output := make([]string, maxLen)

	for i, s := range input {
		for j, c := range s {
			for len(output[j]) < i {
				output[j] += " "
			}
			output[j] += string(c)
		}
	}

	return output
}
