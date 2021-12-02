package main

import (
    "lesson15/lib/table"
    "lesson15/lib/temperature"
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


func drawTable(getTable func() *table.Table) {
    getTable().Display()
}
func GetCelsiusFahrenheitTable() *table.Table {
    return table.GetTemperaturePairTable(&temperature.Celsius{})
}
func GetFahrenheitCelsiusTable() *table.Table {
    return table.GetTemperaturePairTable(&temperature.Fahrenheit{})
}

func main() {
    // 表1
    drawTable(GetCelsiusFahrenheitTable)
    // 表2
    drawTable(GetFahrenheitCelsiusTable)
}
