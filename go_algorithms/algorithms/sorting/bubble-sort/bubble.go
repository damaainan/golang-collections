// Package main provides ...
package bubble

import (
	"fmt"
	"go_algorithms/algorithms/sorting/utils"
)

func sort(arr []int) {
	for itemCount := len(arr) - 1; ; itemCount-- {
		swap := false
		for i := 1; i <= itemCount; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
}

func BubbleSort() {
	list := utils.GetArrayOfSize(100)

	sort(list)

	//for i := 0; i < len(list)-2; i++ {
	//	if list[i] > list[i+1] {
	fmt.Println(list)
	//	}
	//}
}
