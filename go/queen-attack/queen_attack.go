package queenattack

import (
	"errors"
)

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if whitePosition == blackPosition {
		return false, errors.New("same position")
	}
	if len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false, errors.New("invalid position")
	}
	wY := whitePosition[0]
	wX := whitePosition[1]
	bY := blackPosition[0]
	bX := blackPosition[1]

	if wY < 'a' || wY > 'h' || wX < '1' || wX > '8' ||
		bY < 'a' || bY > 'h' || bX < '1' || bX > '8' {
		return false, errors.New("invalid position")
	}

	return wY == bY || wX == bX || (wX+wY == bX+bY || wY-wX == bY-bX), nil
}
