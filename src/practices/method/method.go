package method

import "fmt"

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c - celsius(absoluteTemperature))
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

type kelvin float64

func (k kelvin) celsius() celsius {
	return celsius(k + kelvin(absoluteTemperature))
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

var absoluteTemperature = -273.15

func Answer() {
	var c celsius = 20.0
	fmt.Printf("%f°Cは、%f°Kです。\n", c, c.kelvin())
	fmt.Printf("%f°Cは、%f°Fです。\n", c, c.fahrenheit())

	var f fahrenheit = 20.0
	fmt.Printf("%f°Fは、%f°Cです。\n", f, f.celsius())
	fmt.Printf("%f°Fは、%f°Kです。\n", f, f.kelvin())

	var k kelvin = 20.0
	fmt.Printf("%f°Kは、%f°Cです。\n", k, k.celsius())
	fmt.Printf("%f°Kは、%f°Fです。\n", k, k.fahrenheit())
}
