// Package weather provides tooling for forecasting.
package weather

// CurrentCondition represents the current weather conditions.
var CurrentCondition string
// CurrentLocation represents the current weather location.
var CurrentLocation string

// Forecast returns a stringified response based on city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
