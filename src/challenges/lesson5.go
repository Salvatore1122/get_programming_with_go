package lesson5

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// ======================================================================================================
// LESSON5
// チャレンジ：火星行きのチケット
// 火星への旅を企画する時は、複数の航宙会社のチケットを一覧できたら便利でしょう。
// 航空会社のチケット料金なら、まとめWebサイトが存在しますが、今のところ航宙会社にはないみたいですね。
// でも大丈夫です。Goを使えば、あなたのコンピュータにこういう問題を解かせることが出来るのですから。
// まずはプロトタイプとして10個のチケットをランダムに生成し、きちんと見出しをつけた表形式で次のように表示させます。
//
// Spaceline         Days  Trip type    Price
// ==========================================
// Virgin Galactic   23    Round-trip   $ 96
// Virgin Galactic   29    One-way      $ 37
// SpaceX            31    One-way      $ 41
// Space Adventures  22    Round-trip   $ 100
// Space Adventures  22    One-way      $ 50
// Virgin Galactic   30    Round-trip   $ 84
// Virgin Galactic   24    Round-trip   $ 94
// Space Adventures  27    One-way      $ 44
// Space Adventures  28    Round-trip   $ 86
// SpaceX            41    Round-trip   $ 72
//
// この表には4つの列があります。
// * サービスを提供する航宙会社（Spaceline）
// * 火星への旅（片道）にかかる日数（Days）
// * 旅の種類（Trip type）として、価格が往復（Round-trip）か片道（One-way）か
// * 価格（Price）を百万ドル単位で！
//
// それぞれの航宙会社は、Space Adventures, SpaceX, Virgin Galacticから一つをランダムに選ぶことにします。
// 全てのチケットで出発日には2020年の10月13日を使います。その頃、火星は地球から62,100,000km 離れた位置にあるはずです。
// 宇宙船の飛行速度は、16km/秒から30km/秒の範囲でランダムに選びます。これで旅行の日数が決まり、チケットの料金も速度で決まります。
// 高速な宇宙船ほど高額で、その範囲で百万ドル単位で36から50までとします。往復ならば倍額です。
// ======================================================================================================

var distance = 62_100_000
var spaceLines = [...]string{"Virgin Galactic", "SpaceX", "Space Adventures"}
var tripTypes = [...]string{"Round-trip", "One-way"}
var minSpeed = 16
var maxSpeed = 30
var minPrice = 36

func Answer()  {
	rand.Seed(time.Now().UnixNano())

	outputFormat := "%-18v%-6v%-13v%-5v\n"
	fmt.Printf(outputFormat, "Speceline", "Days", "Trip type", "Price")
	fmt.Println("==========================================")

	spaceLinesLength := len(spaceLines)
	tripTypesLength := len(tripTypes)
	for count := 0; count < 10; count++ {
		speed := minSpeed + rand.Intn(maxSpeed - minSpeed) + 1
		tripType := tripTypes[rand.Intn(tripTypesLength)]
		price := minPrice + (speed - minSpeed)
		if tripType == "Round-trip" {
			price *= 2
		}

		fmt.Printf(
			outputFormat,
			spaceLines[rand.Intn(spaceLinesLength)],
			distance/(speed * 24 * 60 * 60),
			tripType,
			"$ " + strconv.Itoa(price))
	}
}
