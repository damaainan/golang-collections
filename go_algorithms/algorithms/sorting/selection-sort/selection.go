package selection

import (
	"fmt"
	"go_algorithms/algorithms/sorting/utils"
)

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

func SelectionSort() {
	list := utils.GetArrayOfSize(100)

	sort(list)

	fmt.Println(list)

}
