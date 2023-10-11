package convert

import (
	"math"
)

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
	celsius := (fahrenheit - 32) * 5 / 9
	var rounder float64
	pow := math.Pow(10, float64(3))
	intermed := celsius * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}

func TonneToAtomicMass(tonne float64) (result float64, err error) {

	return tonne * 6.02214076e26, nil
}

func KilometerToAngstrom(kilometer float64) (result float64, err error) {

	return kilometer * 1e13, nil
}

func CelsiusToPlanck(celsius float64) (result float64, err error) {

	return celsius * 1.416808 * (1e32), nil
}
