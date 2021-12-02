package temperature

import "errors"

const MIN_CELSIUS_VALUE = -273.15



type CelsiusValue = TemperatureValue
func newCelsiusValue(value float64) (*CelsiusValue, error) {
	if value < MIN_CELSIUS_VALUE {
		return nil, errors.New("摂氏として扱えません。")
	}

	return &value, nil
}


type CelsiusInterface interface {
	TemperatureInterface
}
func NewCelsius(value float64) (CelsiusInterface, error) {
	cv, err := newCelsiusValue(value)
	if err != nil {
		return nil, err
	}

	return &Celsius{value: *cv}, nil
}

type Celsius struct {
	value CelsiusValue
}


func (c *Celsius) Value() CelsiusValue {
	return c.value
}

func (*Celsius) Symbol() Symbol {
	return "°C"
}

func (c *Celsius) ToCelsius() (CelsiusInterface, error) {
	return c, nil
}

func (c *Celsius) ToFahrenheit() (FahrenheitInterface, error) {
	return NewFahrenheit((c.value * 9.0 / 5.0) + 32.0)
}