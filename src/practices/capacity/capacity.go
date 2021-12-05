package capacity

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: 長さ %v 容量 %v\n", label, len(slice), cap(slice))
}

func Answer(loopCount int) {
	capacity := 10
	targetSlice := make([]string, 0, capacity)

	dump("最初", targetSlice)
	for i := 1; i < loopCount; i++ {
		iStr := fmt.Sprintf("%v", i)
		targetSlice = append(targetSlice, iStr)

		if cap(targetSlice) != capacity {
			capacity = cap(targetSlice)
			dump(iStr+"回目のループで容量が変化しました。", targetSlice)
		}
	}
}
