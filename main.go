package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestingTestConverterPoundsToKilogramsNull(t *testing.T) {
	mean := ConvertPoundsToKilograms(0)
	assert.Equal(t, 0.0, mean, "Неверное определение килограмм для нулевого значения")
}
func TestingTestConverterPoundsToKilogramsSimple(t *testing.T) {
	mean := ConvertPoundsToKilograms(22)
	assert.Equal(t, 9.979, mean, "Неверное определение килограмм для значения в фунтах")
}
func TestingTestConverterPoundsToKilogramsMinusVar(t *testing.T) {
	mean := ConvertPoundsToKilograms(-30)
	assert.Equal(t, 0.0, mean, "Неверное определение килограмм для отрицательного значения")
}

func TestingTestConvertFeetToMetersNull(t *testing.T) {
	mean := ConvertFeetToMeters(0)
	assert.Equal(t, 0.0, mean, "Неверное определение метров для нулевого значения")
}
func TestingConvertFeetToMetersSimple(t *testing.T) {
	mean := ConvertFeetToMeters(5)
	assert.Equal(t, 1.524, mean, "Неверное определение метров для значения в футов")
}
func TestingTestConvertFeetToMetersMinusVar(t *testing.T) {
	mean := ConvertFeetToMeters(-30)
	assert.Equal(t, 0.0, mean, "Неверное определение метров для отрицательного значения")
}

func TestingTestConvertFahrenheitToCelsiusNull(t *testing.T) {
	mean := ConvertFahrenheitToCelsius(0)
	assert.Equal(t, -17.778, mean, "Неверное определение цельсий для нулевого значения")
}
func TestingConvertFahrenheitToCelsiusSimple(t *testing.T) {
	mean := ConvertFeetToMeters(112)
	assert.Equal(t, 44.444, mean, "Неверное определение цельсий для значения в фаренгейтах")
}
func TestingTestConvertFahrenheitToCelsiusMinusVar(t *testing.T) {
	mean := ConvertFeetToMeters(-30)
	assert.Equal(t, -34.445, mean, "Неверное определение цельсий для отрицательного значения")
}
