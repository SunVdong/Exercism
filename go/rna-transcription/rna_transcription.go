package strand

var m = map[rune]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	rna := make([]byte, len(dna))
	for i, c := range dna {
		char, ok := m[c]
		if !ok {
			return ""
		}
		rna[i] = char
	}

	return string(rna)
}
