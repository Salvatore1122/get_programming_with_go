package canis

import "fmt"

// ======================================================================================================
// LESSON8
// 練習問題
// おおいぬ座矮小銀河は知られている限り、銀河系に最も近い銀河で、太陽からの距離は236,000,000,000,000,000kmです。
// この距離を定数を使って光年に変換してください。
// ======================================================================================================

const canisDistance = 236_000_000_000_000_000
const lightSpeed = 299792
const secondsPerYear = 60 * 60 * 24 * 365

func Answer() {
	fmt.Printf("太陽からおおいぬ座矮小銀河までの距離は、%v 光年です。\n", canisDistance/lightSpeed/secondsPerYear)
}
