package sort

import (
	"errors"

	"github.com/sanamlimbu/dsa"
	h "github.com/sanamlimbu/dsa/heap"
)

// Heap sort is a comparison-based sorting technique based on Binary Heap Data Structure.
// It can be seen as an optimization over selection sort where we first find the max (or min) element and
// swap it with the last (or first). We repeat the same process for the remaining elements.
// In Heap Sort, we use Binary Heap so that we can quickly find and move the max element in O(Log n)
// instead of O(n) and hence achieve the O(n Log n) time complexity.
//
// Time complexity: O(nlogn)
func HeapSort[T dsa.Ordered](input []T, sortType dsa.SortType) []T {
	copied := make([]T, len(input))
	copy(copied, input)

	result := make([]T, len(input))

	heap := h.NewHeap[T](h.MaxHeapType, copied)

	for i := len(result) - 1; i >= 0; i-- {
		item, err := heap.Extract()
		if err != nil && errors.Is(err, h.ErrHeapEmpty) {
			break
		}

		result[i] = item
	}

	return result
}
