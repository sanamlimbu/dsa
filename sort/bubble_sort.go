package sort

import "github.com/sanamlimbu/dsa"

// Bubble sort is a sorting algorithm that compares two adjacent elements
// and swaps them until they are in the intended order.
//
// It is an in-place algorithm.
//
// Time complexity: O(n^2)
func BubbleSort[T dsa.Ordered](input []T, sortType dsa.SortType) {
	for i := range input {

		for j := 0; j < len(input)-i-1; j++ {

			if sortType == dsa.DescendingSortType {
				if input[j] < input[j+1] {
					input[j], input[j+1] = input[j+1], input[j]
				}

			} else {
				if input[j] > input[j+1] {
					input[j], input[j+1] = input[j+1], input[j]
				}
			}
		}
	}
}
