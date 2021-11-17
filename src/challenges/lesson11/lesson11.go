package lesson11

import (
	"fmt"
	"strings"
)

var cipherText = "CSOITEUIWUIZNSROCNKF"
var originalText = "your message goes here"
var keyword = "GOLANG"

func DecipherAnswer() {
	keywordLength := len(keyword)

	// rangeを使わないパターン
	cipherTextLength := len(cipherText)
	for i := 0; i < cipherTextLength; i++ {
		c := rune(cipherText[i]) - (rune(keyword[i%keywordLength]) - 'A')
		if c < 'A' {
			c += 26
		}
		fmt.Printf("%c", c)
	}

	// rangeを使うパターン
	for i, c := range cipherText {
		c = c - (rune(keyword[i%keywordLength]) - 'A')
		if c < 'A' {
			c += 26
		}
		fmt.Printf("%c", c)
	}
}

func CipherAnswer() {
	keywordLength := len(keyword)
	convertedText := strings.ToUpper(strings.Replace(originalText, " ", "", -1))
	for i, c := range convertedText {
		if c != ' ' {
			c = c + (rune(keyword[i%keywordLength]) - 'A')
			if c > 'Z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
}
