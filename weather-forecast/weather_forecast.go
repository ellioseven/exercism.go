// Package weather is a package.
package weather

// CurrentCondition Weather condition.
var CurrentCondition string

// CurrentLocation Weather location.
var CurrentLocation string

// Forecast displays the weather condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
