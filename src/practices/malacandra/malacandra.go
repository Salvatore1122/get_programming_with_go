package malacandra

import "fmt"

func Answer()  {
	const hoursPerDay = 24
	const distance = 56000000
	var targetDays = 28

	fmt.Println(distance / (targetDays * hoursPerDay), "km/時")
}
