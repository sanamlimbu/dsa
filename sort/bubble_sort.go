package sort

import "github.com/sanamlimbu/dsa"

// Bubble sort is a sorting algorithm that compares two adjacent elements
// and swaps them until they are in the intended order.
//
// Time complexity: O(n^2)
func BubbleSort[T dsa.Ordered](input []T, sortType dsa.SortType) []T {
	result := make([]T, len(input))

	copy(result, input)

	length := len(result)

	for i := range result {
		for j := 0; j < length-i-1; j++ {
			if sortType == dsa.DescendingSortType {
				if result[j] < result[j+1] {
					result[j], result[j+1] = result[j+1], result[j]
				}
			} else {
				if result[j] > result[j+1] {
					result[j], result[j+1] = result[j+1], result[j]
				}
			}
		}
	}

	return result
}
