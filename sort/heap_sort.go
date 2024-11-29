package sort

import (
	"errors"

	"github.com/sanamlimbu/dsa"
	h "github.com/sanamlimbu/dsa/heap"
)

// Heap sort is a comparison-based sorting technique based on binary heap data structure.
// It can be seen as an optimization over selection sort where we first find the max (or min) element and
// swap it with the last (or first). We repeat the same process for the remaining elements.
// In heap sort, we use binary heap so that we can quickly find and move the max element in O(logn)
// instead of O(n) and hence achieve the O(nlogn) time complexity.
//
// It is an in-place algorithm.
//
// Time complexity: O(nlogn)
func HeapSort[T dsa.Ordered](input []T, sortType dsa.SortType) {
	heap := h.NewHeap[T](h.MaxHeapType, input)

	for i := len(input) - 1; i >= 0; i-- {

		item, err := heap.Extract()
		if err != nil && errors.Is(err, h.ErrHeapEmpty) {
			break
		}

		input[i] = item
	}
}
