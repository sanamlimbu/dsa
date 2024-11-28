package sort

import "github.com/sanamlimbu/dsa"

// Selection sort is a comparison-based sorting algorithm.
// It sorts an array by repeatedly selecting the smallest (or largest) element from the unsorted portion
// and swapping it with the first unsorted element. This process continues until the entire array is sorted.
//
// Time complexity: O(n^2)
func SelectionSort[T dsa.Ordered](input []T, sortType dsa.SortType) []T {
	length := len(input)

	result := make([]T, length)
	copy(result, input)

	for i := 0; i < length-1; i++ {
		index := i

		for j := i + 1; j < length; j++ {

			if sortType == dsa.DescendingSortType {
				if result[j] > result[index] {
					index = j
				}
			} else {
				if result[j] < result[index] {
					index = j
				}
			}
		}

		result[i], result[index] = result[index], result[i]
	}

	return result
}
