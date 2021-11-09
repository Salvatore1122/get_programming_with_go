package random_dates

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func Answer()  {
	for count := 10; count > 0; count-- {
		year := rand.Intn(2021) + 1
		month := rand.Intn(12) + 1
		dayInMonth := 31
		switch month {
		case 2:
			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				dayInMonth = 29
			} else {
				dayInMonth = 28
			}
		case 4, 6, 9, 11:
			dayInMonth = 30
		}
		day := rand.Intn(dayInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}