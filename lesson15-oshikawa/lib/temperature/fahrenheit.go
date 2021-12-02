package temperature

import (
	"errors"
)

const MIN_FAHRENHEIT_VALUE = -459.67



type FahrenheitValue = TemperatureValue
func newFahrenheitValue(value float64) (*FahrenheitValue, error) {
	if value < MIN_FAHRENHEIT_VALUE {
		return nil, errors.New("華氏として扱えません。")
	}

	return &value, nil
}


type FahrenheitInterface interface {
	TemperatureInterface
}
func NewFahrenheit(value float64) (FahrenheitInterface, error) {
	fv, err := newFahrenheitValue(value)
	if err != nil {
		return nil, err
	}

	return &Fahrenheit{value: *fv}, nil
}

type Fahrenheit struct {
	value FahrenheitValue
}


func (f *Fahrenheit) Value() FahrenheitValue {
	return f.value
}

func (*Fahrenheit) Symbol() Symbol {
	return "°F"
}

func (f *Fahrenheit) ToCelsius() (CelsiusInterface, error) {
	return NewCelsius((f.value - 32.0) * 5.0 / 9.0)
}

func (f *Fahrenheit) ToFahrenheit() (FahrenheitInterface, error) {
	return f, nil
}