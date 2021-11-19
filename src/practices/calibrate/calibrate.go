package calibrate

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func Answer() {
	var offset kelvin = 10
	realSensorFn := calibrate(realSensor, offset)
	// ここでoffsetの値を変えてもsensor関数にはoffset = 10がセットされている
	offset++
	fmt.Println(realSensorFn())

	// fakeSensor内で呼ばれている
	fakeSensorFn := calibrate(fakeSensor, offset)
	for i := 0; i < 10; i++ {
		fmt.Println(fakeSensorFn())
		time.Sleep(time.Second)
	}
}
