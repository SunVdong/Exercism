package resistorcolortrio

import "fmt"

var colorCodes = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	// Get the first two colors' values
	value1 := colorCodes[colors[0]]
	value2 := colorCodes[colors[1]]

	// Get the third color's value for the multiplier
	multiplier := colorCodes[colors[2]]

	// Combine the first two values and multiply by 10^multiplier
	resistorValue := (value1*10 + value2) * intPow(10, multiplier)

	// Determine the suffix (ohms or kiloohms)
	var result string
	if resistorValue >= 1_000_000_000 {
		result = fmt.Sprintf("%d gigaohms", resistorValue/1000/1000/1000)
	} else if resistorValue >= 1_000_000 {
		result = fmt.Sprintf("%d megaohms", resistorValue/1000/1000)
	} else if resistorValue >= 1000 {
		result = fmt.Sprintf("%d kiloohms", resistorValue/1000)
	} else {
		result = fmt.Sprintf("%d ohms", resistorValue)
	}

	return result
}

func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}
