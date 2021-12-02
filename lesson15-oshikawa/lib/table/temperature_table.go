package table

import (
	"strconv"

	"lesson15/lib/temperature"
)

func getTemperaturePairTable(
	getCol1 temperature.NewTemperature,
	getCol2 temperature.NewTemperature,
	getCol2FromTemp func (col1Temp temperature.TemperatureInterface) (temperature.TemperatureInterface, error),
) *Table {
	headerCol1, _ := getCol1(0)
	headerCol2, _ := getCol2(0)

	header := [2]string{headerCol1.Symbol(), headerCol2.Symbol()}
	var body [][2]string

	for temp := -40.0; temp <= 100.0; temp+=5.0 {
		bodyCol1, _ := getCol1(temp)
		bodyCol2, _ := getCol2FromTemp(bodyCol1)
		body = append(
			body,
			[2]string{
				strconv.FormatFloat(bodyCol1.Value(), 'f', 2, 64),
				strconv.FormatFloat(bodyCol2.Value(), 'f', 2, 64),
			},
		)
	}

	return NewTable(
		header,
		body,
	)
}


func GetTemperaturePairTable(firstColTemp temperature.TemperatureInterface) *Table {
	switch firstColTemp.(type) {
	case *temperature.Celsius:
		return getTemperaturePairTable(
			temperature.NewCelsiusAsTemp,
			temperature.NewFahrenheitAsTemp,
			func (col1Temp temperature.TemperatureInterface) (temperature.TemperatureInterface, error) {
				return col1Temp.ToFahrenheit()
			},
		)
	case *temperature.Fahrenheit:
		return getTemperaturePairTable(
			temperature.NewFahrenheitAsTemp,
			temperature.NewCelsiusAsTemp,
			func (col1Temp temperature.TemperatureInterface) (temperature.TemperatureInterface, error) {
				return col1Temp.ToCelsius()
			},
		)
	default:
		return nil
	}
}