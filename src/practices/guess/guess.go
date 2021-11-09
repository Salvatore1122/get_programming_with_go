package guess

import (
	"fmt"
	"math/rand"
)

func Answer()  {
	const correctNumber = 67

	label1:
		for {
			var num = rand.Intn(100) + 1

			switch {
			case num < correctNumber:
				fmt.Printf("%v: 小さすぎます！\n", num)
			case num > correctNumber:
				fmt.Printf("%v: 大きすぎます！\n", num)
			default:
				fmt.Printf("%v: 正解！\n", num)
				break label1
			}
		}

	// 解答例
	//var number = 42
	//for {
	//	var n = rand.Intn(100) + 1
	//	if n < number {
	//		fmt.Printf("%v is too small.\n", n)
	//	} else if n > number {
	//		fmt.Printf("%v is too big.\n", n)
	//	} else {
	//		fmt.Printf("You got it! %v\n", n)
	//		break
	//	}
	//}
}