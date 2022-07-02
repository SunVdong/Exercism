package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// SillyNephewError Define the SillyNephewError type here.
type SillyNephewError struct {
	Num int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.Num)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction {
			fodder *= 2
		} else {
			return 0, errors.New("non-scale error")
		}
	}

	if fodder < 0 {
		return 0, errors.New("negative fodder")
	}

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0, &SillyNephewError{Num: cows}
	}

	return fodder / float64(cows), nil
}
