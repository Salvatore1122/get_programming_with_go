package international

import "fmt"

// ======================================================================================================
// LESSON9
// 練習問題-2
// スペイン語のメッセージ「Hola Estación Espacial Internacional」をROT13で暗号化してください。
// rangeキーワードを使えばスペイン語のテキストにROT13を使っても、アクセント付きの文字が保存されるでしょう。
// ======================================================================================================

func Answer() {
	message := "Hola Estación Espacial Internacional"
	for _, c := range message {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
			c += 13
			if c > 26 {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
}
