package bs

import (
	"fmt"
)

func search(sortedArray []int, el int) int {
	init, end := 0, len(sortedArray)-1

	for init <= end {
		middle := ((end - init) >> 1) + init

		if sortedArray[middle] == el {
			return middle
		}

		if sortedArray[middle] < el {
			init = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}

func Bs() {
	sorted := make([]int, 10000)

	for i := 0; i < 10000; i++ {
		sorted[i] = 2 * i
	}

	for i := 0; i < 10000; i++ {
		index := search(sorted, 2*i)

		if index == i {
			fmt.Println(index)
		}
	}

	if search(sorted, 3) != -1 {
		fmt.Println(11)
	}
}
