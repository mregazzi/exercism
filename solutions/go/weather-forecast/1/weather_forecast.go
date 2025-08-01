// Package weather provides tools for forecasting weather of various cities in Goblonocus.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string
// CurrentLocation represents the city the condition refers to.
var CurrentLocation string

// Forecast returns the current weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
