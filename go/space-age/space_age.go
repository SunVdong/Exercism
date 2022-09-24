package space

import (
	"strings"
)

type Planet string

func Age(seconds float64, planet Planet) float64 {
	year_earth := seconds / 31557600
	switch strings.ToLower(string(planet)) {
	case "mercury":
		return year_earth / 0.2408467
	case "venus":
		return year_earth / 0.61519726
	case "earth":
		return year_earth
	case "mars":
		return year_earth / 1.8808158
	case "jupiter":
		return year_earth / 11.862615
	case "saturn":
		return year_earth / 29.447498
	case "uranus":
		return year_earth / 84.016846
	case "neptune":
		return year_earth / 164.79132
	default:
		return -1
	}
}
