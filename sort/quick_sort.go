package sort

import (
	"math/rand"
	"time"

	"github.com/sanamlimbu/dsa"
)

// QuickSort works on the principle of divide and conquer.
// There are mainly three steps in the algorithm:
//
//  1. Choose a pivot: Select an element from the array as the pivot. The choice of pivot can vary
//     (e.g., first element, last element, random element, or median).
//
//  2. Partition the array: Rearrange the array around the pivot. After partitioning,
//     all elements smaller than the pivot will be on its left, and all elements greater than the pivot will be on its right.
//     The pivot is then in its correct position, and we obtain the index of the pivot.
//
//  3. Recursively call: Recursively apply the same process to the two partitioned sub-arrays (left and right of the pivot).
//
//  4. Base case: The recursion stops when there is only one element left in the sub-array, as a single element is already sorted.
//
// Complexity: Worst case O(n^2) , Average case O(n log n)
func QuickSort[T dsa.Ordered](input []T, sortType dsa.SortType) []T {
	length := len(input)

	if length < 2 {
		return input
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(length)

	pivot := input[index]

	less := make([]T, 0)
	greater := make([]T, 0)

	for i, item := range input {
		if i == index {
			continue
		}

		if item <= pivot {
			less = append(less, item)
		} else {
			greater = append(greater, item)
		}
	}

	less = QuickSort(less, sortType)
	greater = QuickSort(greater, sortType)

	if sortType == dsa.DescendingSortType {
		return append(append(greater, pivot), less...)
	}

	return append(append(less, pivot), greater...)
}
