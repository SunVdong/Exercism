// Package weather provides forecasts.
package weather

// CurrentCondition mean current weather condition.
var CurrentCondition string
// CurrentLocation mean current location.
var CurrentLocation string

// Forecast return current weather condition of some city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
