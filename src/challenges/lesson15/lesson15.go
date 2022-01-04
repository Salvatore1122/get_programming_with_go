package lesson15

import "fmt"

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

var (
	minDisplayVal      = -40
	intervalDisplayVal = 5
)

type celsius float64
type fahrenheit float64
type getRowFn func(row int) (string, string)

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func getCelsiusAndFahrenheit(row int) (string, string) {
	c := celsius(row*intervalDisplayVal + minDisplayVal)
	return fmt.Sprintf("%.2f", c), fmt.Sprintf("%.2f", c.fahrenheit())
}

func getFahrenheitAndCelsius(row int) (string, string) {
	f := fahrenheit(row*intervalDisplayVal + minDisplayVal)
	return fmt.Sprintf("%.2f", f), fmt.Sprintf("%.2f", f.celsius())
}

func drawTable(hdr1, hdr2 string, rowCount int, getRow getRowFn) {
	format := "| %-8s | %-8s |\n"
	line := "======================="

	fmt.Println(line)
	fmt.Printf(format, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rowCount; row++ {
		cel1, cel2 := getRow(row)
		fmt.Printf(format, cel1, cel2)
	}
	fmt.Println(line)
}

func Answer() {
	// 表1
	drawTable("°C", "°F", 29, getCelsiusAndFahrenheit)

	// 表2
	drawTable("°F", "°C", 29, getFahrenheitAndCelsius)
}
