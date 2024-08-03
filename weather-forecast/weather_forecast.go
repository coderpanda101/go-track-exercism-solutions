// Package weather contains functions and variables to forecast weather condition.
package weather

// CurrentCondition specifies current weather condition.
var CurrentCondition string

// CurrentLocation specifies current location.
var CurrentLocation string

// Forecasts weather based on condition and city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
