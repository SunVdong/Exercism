package perfect

import "fmt"

type Classification int8

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = fmt.Errorf("Cannot process negative number")

func Classify(num int64) (Classification, error) {
	if num <= 0 {
		return -1, ErrOnlyPositive
	}
	if num == 1 {
		return ClassificationDeficient, nil
	}
	sum := int64(1)
	for i := int64(2); i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if num/i != i {
				sum += num / i
			}
		}
	}
	if sum == num {
		return ClassificationPerfect, nil
	} else if sum > num {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}
}
