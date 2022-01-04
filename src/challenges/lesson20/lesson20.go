package lesson20

import "fmt"

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func (u Universe) show() {
	for _, row := range u {
		for _, col := range row {
			if col {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func newUniverse() Universe {
	univ := make(Universe, height, height)
	for i := 0; i < height; i++ {
		univ[i] = make([]bool, width, width)
	}
	return univ
}

func Answer() {
	newUniverse().show()
}
