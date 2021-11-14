package input

import (
	"fmt"
	"math/rand"
)

// ======================================================================================================
// LESSON10
// 練習問題
// 文字列をbool型に変換するプログラムを書きましょう。
// * 文字列の「true」「yes」「1」はどれも true にします。
// * 文字列の「false」「no」「0」はどれも false にします。
// * 他の値ならエラーメッセージを表示してください。
// ======================================================================================================

var strArray = [...]string{"true", "false", "yes", "no", "1", "0"}

func Answer() {
	strArrayLength := len(strArray)
	targetStr := strArray[rand.Intn(strArrayLength)]

	switch targetStr {
	case "true", "yes", "1":
		fmt.Println(true)
	case "false", "no", "0":
		fmt.Println(false)
	default:
		fmt.Println("bool型への変換でエラーが発生しました。")
	}
}
