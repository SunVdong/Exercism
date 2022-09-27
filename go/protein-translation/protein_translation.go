package protein

import (
	"errors"
	"fmt"
)

func FromRNA(rna string) ([]string, error) {
	l := len(rna)

	if l <= 0 {
		return nil, fmt.Errorf("too short")
	}

	if l%3 != 0 {
		return nil, fmt.Errorf("invalid length: %s", rna)
	}

	condons := ""
	var res []string
	for i, v := range rna {
		condons += string(v)
		if i%3 == 2 {
			res = append(res, condons)
			condons = ""
		}
	}

	return res, nil
}

func FromCodon(codon string) (string, error) {
	var str string
	switch codon {
	case "AUG":
		str += "Methionine"
	case "UUU", "UUC":
		str += "Phenylalanine"
	case "UUA", "UUG":
		str += "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		str += "Serine"
	case "UAU", "UAC":
		str += "Tyrosine"
	case "UGU", "UGC":
		str += "Cysteine"
	case "UGG":
		str += "Tryptophan"
	case "UAA", "UAG", "UGA":
		// return "", fmt.Errorf("Stop codon: %s", codon)
		return "", errors.New("Stop codon")
	default:
		return "", fmt.Errorf("Unknown codon: %s", codon)
	}

	return str, nil
}
