package temperature


type TemperatureValue = float64
type Symbol = string

type TemperatureInterface interface {
	Value() TemperatureValue
	Symbol() Symbol
	ToCelsius() (CelsiusInterface, error)
	ToFahrenheit() (FahrenheitInterface, error)
}

type NewTemperature func (value float64) (TemperatureInterface, error)

func NewCelsiusAsTemp(value float64) (TemperatureInterface, error) {
	return NewCelsius(value)
}
func NewFahrenheitAsTemp(value float64) (TemperatureInterface, error) {
	return NewFahrenheit(value)
}