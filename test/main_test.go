package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConverterPoundsToKilogramsNull(t *testing.T) {
	mean := ConvertPoundsToKilograms(0)
	assert.Equal(t, 0.0, mean, "Неверное определение килограмм для нулевого значения")
}
func TestConverterPoundsToKilogramsSimple(t *testing.T) {
	mean := ConvertPoundsToKilograms(22)
	assert.Equal(t, 9.979, mean, "Неверное определение килограмм для значения в фунтах")
}
func TestConverterPoundsToKilogramsMinusVar(t *testing.T) {
	mean := ConvertPoundsToKilograms(-30)
	assert.Equal(t, 0.0, mean, "Неверное определение килограмм для отрицательного значения")
}

func TestTestConvertFeetToMetersNull(t *testing.T) {
	mean := ConvertFeetToMeters(0)
	assert.Equal(t, 0.0, mean, "Неверное определение метров для нулевого значения")
}
func TestConvertFeetToMetersSimple(t *testing.T) {
	mean := ConvertFeetToMeters(5)
	assert.Equal(t, 1.524, mean, "Неверное определение метров для значения в футов")
}
func TestTestConvertFeetToMetersMinusVar(t *testing.T) {
	mean := ConvertFeetToMeters(-30)
	assert.Equal(t, 0.0, mean, "Неверное определение метров для отрицательного значения")
}

func TestConvertFahrenheitToCelsiusNull(t *testing.T) {
	mean := ConvertFahrenheitToCelsius(0)
	assert.Equal(t, -17.778, mean, "Неверное определение цельсий для нулевого значения")
}
func TestConvertFahrenheitToCelsiusSimple(t *testing.T) {
	mean := ConvertFeetToMeters(112)
	assert.Equal(t, 44.444, mean, "Неправильное определение цельсий для значения в фаренгейтах")
}
func TestTestConvertFahrenheitToCelsiusMinusVar(t *testing.T) {
	mean := ConvertFeetToMeters(-30)
	assert.Equal(t, -34.445, mean, "Неверное определение цельсий для отрицательного значения")
}
