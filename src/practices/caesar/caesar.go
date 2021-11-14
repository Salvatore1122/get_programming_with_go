package caesar

import "fmt"

// ======================================================================================================
// LESSON9
// 練習問題-1
// ユリウス・カエサルの言葉とされる、次の暗号を解きましょう。
// L fdph, L vdz, L frqtxhuhg
// プログラムでは、大文字と小文字をそれぞれ3文字分ずつシフトして戻す必要があります。
// ただし、単純に3を引くのではなく、「a」が「x」に、「b」が「y」に、「c」が「z」になるように工夫する必要があります。
// ======================================================================================================

func Answer() {
	message := "L fdph, L vdz, L frqtxhuhg"
	for _, c := range message {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
			c -= 3
			if c < 0 {
				c += 26
			}
		}

		fmt.Printf("%c", c)
	}
}