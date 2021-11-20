package chess

import (
	"fmt"
)

func displayBoard(board [8][8]rune) {
	for _, row := range board {
		for _, c := range row {
			if c == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
}

func Answer() {
	var board [8][8]rune
	var kqrbn = [...]rune{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'}
	var KQRBN = [...]rune{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'}

	// ルーク・ナイト・ビショップ・キング・クイーン
	for _, i := range [...]int{0, 7} {
		for j := range board[i] {
			if i == 0 {
				board[i][j] = kqrbn[j]
			} else {
				board[i][j] = KQRBN[j]
			}
		}
	}

	// ポーン
	for _, i := range [...]int{1, 6} {
		for j := range board[i] {
			if i == 1 {
				board[i][j] = 'p'
			} else {
				board[i][j] = 'P'
			}
		}
	}

	displayBoard(board)
}
