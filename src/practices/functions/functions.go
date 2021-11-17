package functions

import "fmt"

const absoluteTemperature = -273.15

func Answer() {
	// 233°K → ?°C
	fmt.Printf("233°Kは %f °Cです。", kelvinToCelsius(233.0))

	// 0°K → ?°F
	fmt.Printf("0°Kは %f °Fです。", kelvinToFahrenheit(0.0))
}

func kelvinToCelsius(k float64) float64 {
	return k + absoluteTemperature
}

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}
