package lesson15

import (
	"fmt"
	"strings"
)

// ======================================================================================================
// LESSON15
// チャレンジ：温度テーブル
// 温度変換のテーブルを表示するプログラムを書きましょう。
// その表では次のように等号(=)とパイプ(|)の記号を使って線を引き、ヘッダセクションを作ります。
//
// =====================
// | °C      | °F      |
// =====================
// | -40.0   | -40.0   |
// | ...     | ...     |
// =====================
//
// プログラムで、2つの表を作ってください。第1の表は列が2つあり、第1列で°Cを、第2列で°Fを示します。
// -40°Cから100°Cまで、5°Cずつ変化させるループを使います。その中で温度変換メソッドを作り、両方の列に記入します。
// 第1の表が完成したら第2の表を実装し、今度は列を入れ換えます（°Fから°Cへの変換表です）。
// 線を引いたり値をパディングしたりするコードは、2列の表にする必要のあるデータなら、どんなものにも再利用できます。
// 表の線を引くコードと、それぞれの列の温度を計算するコードを関数を使って切り分けましょう。
// drawTable関数を実装しましょう。
// これは1個の第一級関数をパラメータとして受け取り、それを呼び出すことのよって各列に表示するデータを取得します。
// 別の関数をdrawTableに渡すとその結果として別のデータが表示されるようにしてください。
// ======================================================================================================

const (
	min           = -40
	max           = 100
	step          = 5
	lineFormat    = "| %-8s | %-8s |"
	lineSeparator = "======================="
)

type (
	temperatureInterface interface {
		unitLabel() string
		toCelsius() celsiusTemperatureValue
		toFahrenheit() fahrenheitTemperatureValue
	}
	temperatureValue           float64
	celsiusTemperatureValue    temperatureValue
	fahrenheitTemperatureValue temperatureValue
	getTableDataFn             func() tableData
	celsius                    struct {
		temperatureInterface
		value celsiusTemperatureValue
	}
	fahrenheit struct {
		temperatureInterface
		value fahrenheitTemperatureValue
	}
	tableData struct {
		header [2]string
		body   [][2]string
	}
)

func (temperature celsius) unitLabel() string {
	return "°C"
}
func (temperature fahrenheit) unitLabel() string {
	return "°F"
}

func (temperature fahrenheit) toCelsius() celsiusTemperatureValue {
	return celsiusTemperatureValue((temperature.value - 32.0) * 5.0 / 9.0)
}
func (temperature celsius) toFahrenheit() fahrenheitTemperatureValue {
	return fahrenheitTemperatureValue((temperature.value * 9.0 / 5.0) + 32.0)
}

func (temperature celsius) toCelsius() celsiusTemperatureValue {
	return temperature.value
}
func (temperature fahrenheit) toFahrenheit() fahrenheitTemperatureValue {
	return temperature.value
}

func getCelsiusAndFahrenheit() tableData {
	left := celsius{}
	right := fahrenheit{}
	data := tableData{
		header: [2]string{left.unitLabel(), right.unitLabel()},
		body:   [][2]string{},
	}
	for i := min; i <= max; i += step {
		left.value = celsiusTemperatureValue(i)
		data.body = append(
			data.body,
			[2]string{fmt.Sprintf("%.2f", left.toCelsius()), fmt.Sprintf("%.2f", left.toFahrenheit())},
		)
	}
	return data
}

func getFahrenheitAndCelsius() tableData {
	left := fahrenheit{}
	right := celsius{}
	data := tableData{
		header: [2]string{left.unitLabel(), right.unitLabel()},
		body:   [][2]string{},
	}
	for i := min; i <= max; i += step {
		left.value = fahrenheitTemperatureValue(i)
		data.body = append(
			data.body,
			[2]string{fmt.Sprintf("%.2f", left.toFahrenheit()), fmt.Sprintf("%.2f", left.toCelsius())},
		)
	}
	return data
}

func drawTable(getTableData getTableDataFn) {
	table := []string{}
	data := getTableData()

	table = append(table, lineSeparator)
	table = append(table, fmt.Sprintf(lineFormat, data.header[0], data.header[1]))
	table = append(table, lineSeparator)
	for _, row := range data.body {
		table = append(table, fmt.Sprintf(lineFormat, row[0], row[1]))
	}
	table = append(table, lineSeparator)
	fmt.Println(strings.Join(table, "\n"))
}

func Answer() {
	// 表1
	drawTable(getCelsiusAndFahrenheit)

	// 表2
	drawTable(getFahrenheitAndCelsius)
}
