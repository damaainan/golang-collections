package heap

import (
	"fmt"
	"go_algorithms/algorithms/sorting/utils"
	"go_algorithms/data-structures/heap"
)

func sort(arr []int) []int {
	h := heap.NewMin()
	for i := 0; i < len(arr); i++ {
		h.Insert(heap.Int(arr[i]))
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = int(h.Extract().(heap.Int))
	}

	return arr
}

func HeapSort() {
	list := utils.GetArrayOfSize(10)

	sort(list)

	fmt.Println(list)

}
