package insertion

import (
	"fmt"
	"go_algorithms/algorithms/sorting/utils"
)

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > value {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = value
	}
}

func InsertionSort() {
	list := utils.GetArrayOfSize(100)

	sort(list)

	fmt.Println(list)

}
