package protein

import (
	"errors"
)

var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	condon := ""
	var res []string
	for i, v := range rna {
		condon += string(v)
		if i%3 == 2 {
			p, e := FromCodon(condon)
			switch e {
			case ErrStop:
				return res, nil
			case ErrInvalidBase:
				return nil, e
			default:
				res = append(res, p)
				condon = ""
			}
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
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}

	return str, nil
}
