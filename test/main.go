package main

import "math"

func ConvertPoundsToKilograms(pounds float64) float64 {
	kilograms := pounds * 0.45359237
	if pounds < 0 {
		kilograms = 0.0
	}
	var rounder float64
	pow := math.Pow(10, float64(3))
	intermed := kilograms * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
func ConvertFeetToMeters(feet float64) float64 {
	meters := feet * 0.3048
	if feet < 0 {
		meters = 0.0
	}
	return meters
}
func ConvertFahrenheitToCelsius(fahrenheit float64) float64 {
	return 0
}
