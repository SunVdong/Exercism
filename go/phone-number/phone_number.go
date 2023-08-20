package phonenumber

import (
	"errors"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	digits := ""
	for _, c := range phoneNumber {
		if unicode.IsDigit(c) {
			digits += string(c)
		}
	}

	if len(digits) == 10 && (digits[0] >= '2' && digits[0] <= '9') && (digits[3] >= '2' && digits[3] <= '9') {
		return digits, nil
	} else if len(digits) == 11 && digits[0] == '1' && (digits[1] >= '2' && digits[1] <= '9') && (digits[4] >= '2' && digits[4] <= '9') {
		return digits[1:], nil
	} else {
		return "", errors.New("invalid phone number")
	}
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", errors.New("invalid phone number")
	}

	return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", errors.New("invalid phone number")
	}

	// (NXX) NXX-XXXX
	return "(" + string(number[0:3]) + ") " + string(number[3:6]) + "-" + string(number[6:]), nil
}
