package romannumerals

import (
	"bytes"
	"fmt"
)

type arabicToRoman struct {
	arabic int
	roman  string
}

var dictionary = []arabicToRoman{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input <= 0 {
		return "", fmt.Errorf("input must be between 3999 and 0 , now is %d", input)
	}

	buffer := bytes.NewBufferString("")
	for _, item := range dictionary {
		for input >= item.arabic {
			buffer.WriteString(item.roman)
			input -= item.arabic
		}
	}

	return buffer.String(), nil
}
