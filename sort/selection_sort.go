package sort

import "github.com/sanamlimbu/dsa"

// Selection sort is a comparison-based sorting algorithm.
// It sorts an array by repeatedly selecting the smallest (or largest) element from the unsorted portion
// and swapping it with the first unsorted element. This process continues until the entire array is sorted.
//
// It is an in-place algorithm.
//
// Time complexity: O(n^2)
func SelectionSort[T dsa.Ordered](input []T, sortType dsa.SortType) {
	length := len(input)

	for i := 0; i < length-1; i++ {
		index := i

		for j := i + 1; j < length; j++ {

			if sortType == dsa.DescendingSortType {
				if input[j] > input[index] {
					index = j
				}
			} else {
				if input[j] < input[index] {
					index = j
				}
			}
		}

		input[i], input[index] = input[index], input[i]
	}
}
