// Package weather provides functions to forecast the weather.
package weather

// CurrentCondition is a string variable that denotes the current weather condition.
var CurrentCondition string

// CurrentLocation is a string variable that denotes the current location.
var CurrentLocation string

// Forecast returns the weather forecast at a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
