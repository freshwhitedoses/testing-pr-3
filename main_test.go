package testing_pr_3

import (
	"errors"
	"flag"
	_ "flag"
	"os"
	"testing"
	"testing3/convert"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
)

type StepsHandler struct {
	value float64
	error error
}

var handler StepsHandler

var opt = godog.Options{
	Format: "progress",
	Paths:  []string{"features"},
	Output: colors.Colored(os.Stdout),
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()
	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
		Options:             &opt,
	}.Run()
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
func init() {
	godog.BindFlags("godog", flag.CommandLine, &opt)
}
func InitializeScenario(context *godog.ScenarioContext) {
	context.Step(`^i want to convert (-?\d+\.\d+) pound into kilograms`, iWantToConvertPoundIntoKilograms)
	context.Step(`^i want to convert (-?\d+\.\d+) feet into meters`, iWantToConvertFeetIntoMeters)
	context.Step(`^i want to convert (-?\d+\.\d+) fahrenheit to celsius`, iWantToConvertFahrenheitToCelsius)
	context.Step(`^I don't expect error$`, IDontExpectError)
	context.Step(`^I expect error$`, IExpectError)
	context.Step(`^a tonne value of (\d+)$`, aTonneValueOf)
	context.Step(`^I convert it to atomic mass$`, iConvertItToAtomicMass)
	context.Step(`^the result should be (\d+(\.\d+)?)$`, theResultShouldBe)

	context.Step(`^a kilometer value of (\d+)$`, aKilometerValueOf)
	context.Step(`^I convert it to angstrom$`, iConvertItToAngstrom)

	context.Step(`^a celsius value of (\d+)$`, aCelsiusValueOf)
	context.Step(`^I convert it to planck$`, iConvertItToPlanck)
}
func IExpectError() error {
	if handler.error == nil {
		return errors.New("Неверное вычисление")
	}
	return nil
}
func IDontExpectError() error {
	if handler.error != nil {
		return errors.New("Неверное вычисление")
	}
	return nil
}
func iWantToConvertPoundIntoKilograms(pound float64) error {
	answerNew := ConvertPoundsToKilograms(pound)
	if pound < 0 {
		handler.value, handler.error = answerNew, errors.New("Неверное вычисление")
	} else if answerNew < 0 {
		handler.value, handler.error = answerNew, errors.New("Неверное вычисление")
	} else {
		handler.value, handler.error = answerNew, nil
	}
	return nil
}
func iWantToConvertFahrenheitToCelsius(pound float64) error {
	answerNew := ConvertFahrenheitToCelsius(pound)
	if pound < 32 || answerNew >= 0 {
		handler.value, handler.error = answerNew, errors.New("Неверное вычисление")
	}
	handler.value, handler.error = answerNew, nil
	return nil
}
func iWantToConvertFeetIntoMeters(pound float64) error {
	answerNew := ConvertFeetToMeters(pound)
	if pound < 0 {
		handler.value, handler.error = answerNew, errors.New("Неверное вычисление")
	} else if answerNew < 0 {
		handler.value, handler.error = answerNew, errors.New("Неверное вычисление")
	} else {
		handler.value, handler.error = answerNew, nil
	}
	return nil
}

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
	mean := ConvertFahrenheitToCelsius(112)
	assert.Equal(t, 44.444, mean, "Неправильное определение цельсий для значения в фаренгейтах")
}
func TestTestConvertFahrenheitToCelsiusMinusVar(t *testing.T) {
	mean := ConvertFahrenheitToCelsius(-30)
	assert.Equal(t, -34.445, mean, "Неверное определение цельсий для отрицательного значения")
}

func TestTonneToAtomicMass(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
		err      error
	}{
		{1, 6.02214076e26, nil},
		{0, 0, nil},
		{-1, 0, errors.New("Input value must be non-negative")},
	}

	for _, test := range tests {
		result, err := convert.TonneToAtomicMass(test.input)
		if err != nil && test.err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("Expected error message '%s', but got '%s'", test.err.Error(), err.Error())
			}
		} else {
			if result != test.expected {
				t.Errorf("Expected %f, but got %f", test.expected, result)
			}
		}
	}
}

func TestKilometerToAngstrom(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
		err      error
	}{
		{1, 1e13, nil},
		{0, 0, nil},
		{-1, 0, errors.New("Input value must be non-negative")},
	}

	for _, test := range tests {
		result, err := convert.KilometerToAngstrom(test.input)
		if err != nil && test.err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("Expected error message '%s', but got '%s'", test.err.Error(), err.Error())
			}
		} else {
			if result != test.expected {
				t.Errorf("Expected %f, but got %f", test.expected, result)
			}
		}
	}
}

func TestCelsiusToPlanck(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
		err      error
	}{
		{1, 1.416808e32, nil},
		{0, 0, nil},
		{-300, 0, errors.New("Temperature cannot be below absolute zero")},
	}

	for _, test := range tests {
		result, err := convert.CelsiusToPlanck(test.input)
		if err != nil && test.err != nil {
			if err.Error() != test.err.Error() {
				t.Errorf("Expected error message '%s', but got '%s'", test.err.Error(), err.Error())
			}
		} else {
			if result != test.expected {
				t.Errorf("Expected %f, but got %f", test.expected, result)
			}
		}
	}
}

var tonne, kilometer, celsius float64
var result float64

func aTonneValueOf(arg1 float64) error {
	tonne = arg1
	return nil
}

func iConvertItToAtomicMass() error {
	result, _ = convert.TonneToAtomicMass(tonne)
	return nil
}

func theResultShouldBe(arg1 float64) error {
	if arg1 != result {
		return errors.New("RESULT FAILED")
	}
	return nil
}

func aKilometerValueOf(arg1 float64) error {
	kilometer = arg1
	return nil
}

func iConvertItToAngstrom() error {
	result, _ = convert.KilometerToAngstrom(kilometer)
	return nil
}

func aCelsiusValueOf(arg1 float64) error {
	celsius = arg1
	return nil
}

func iConvertItToPlanck() error {
	result, _ = convert.CelsiusToPlanck(celsius)
	return nil
}
