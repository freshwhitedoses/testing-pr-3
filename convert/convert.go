package convert

import (
	"errors"
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

const (
	AvogadroConstant          = 6.02214076e26
	KilometerToAngstromFactor = 1e13
	CelsiusToPlanckFactor     = 1.416808 * (1e32)
	AbsoluteZeroCelsius       = -273.15
)

func TonneToAtomicMass(tonne float64) (float64, error) {
	if math.IsNaN(tonne) || math.IsInf(tonne, 0) {
		return 0, errors.New("Input value must be a finite number")
	}
	if tonne < 0 {
		return 0, errors.New("Input value must be non-negative")
	}
	return tonne * AvogadroConstant, nil
}

func KilometerToAngstrom(kilometer float64) (float64, error) {
	if math.IsNaN(kilometer) || math.IsInf(kilometer, 0) {
		return 0, errors.New("Input value must be a finite number")
	}
	if kilometer < 0 {
		return 0, errors.New("Input value must be non-negative")
	}
	return kilometer * KilometerToAngstromFactor, nil
}

func CelsiusToPlanck(celsius float64) (float64, error) {
	if math.IsNaN(celsius) || math.IsInf(celsius, 0) {
		return 0, errors.New("Input value must be a finite number")
	}
	if celsius < AbsoluteZeroCelsius {
		return 0, errors.New("Temperature cannot be below absolute zero")
	}
	return celsius * CelsiusToPlanckFactor, nil
}
