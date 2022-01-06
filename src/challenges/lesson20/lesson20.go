package lesson20

import (
	"fmt"
	"math/rand"
)

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

func (u Universe) seed() {
	for i, row := range u {
		for j, _ := range row {
			u[i][j] = rand.Intn(2) == 1
		}
	}
}

func (u Universe) alive(x, y int) bool {
	return u[(x+width)%(width-1)][(y+height)%(height-1)]
}

func (u Universe) neighbors(x, y int) int {
	aliveCount := 0
	for _, nums := range [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}} {
		if u.alive(x+nums[0], y+nums[1]) {
			aliveCount++
		}
	}
	return aliveCount
}

func (u Universe) next(x, y int) bool {
	aliveCount := u.neighbors(x, y)
	return aliveCount == 3 || (u.alive(x, y) && aliveCount == 2)
}

func newUniverse() Universe {
	univ := make(Universe, height, height)
	for i := 0; i < height; i++ {
		univ[i] = make([]bool, width, width)
	}
	return univ
}

func Answer() {
	univ := newUniverse()
	univ.seed()
	univ.show()
}
