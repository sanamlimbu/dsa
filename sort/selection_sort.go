package sort

import "github.com/sanamlimbu/dsa"

// Selection sort is a comparison-based sorting algorithm.
// It sorts an array by repeatedly selecting the smallest (or largest) element from the unsorted portion
// and swapping it with the first unsorted element. This process continues until the entire array is sorted.
//
// Time complexity: O(n^2)
func SelectionSort[T dsa.Ordered](input []T, sortType dsa.SortType) []T {
	result := make([]T, 0, len(input))

	for i := range input {
		var j int

		if sortType == dsa.DescendingSortType {
			j = findGreatest(input)
		} else {
			j = findSmallest(input)
		}

		result = append(result, input[j])

		// swap positions
		if i != j {
			input[i], input[j] = input[j], input[i]
		}
	}

	return result
}

// findSmallest returns index of smallest element in given input.
func findSmallest[T dsa.Ordered](input []T) int {
	smallest := input[0]
	index := 0

	for i := 1; i < len(input); i++ {
		if input[i] < smallest {
			smallest = input[i]
			index = i
		}
	}

	return index
}

// findGreatest returns index of greatest element in given input.
func findGreatest[T dsa.Ordered](input []T) int {
	greatest := input[0]
	index := 0

	for i := 1; i < len(input); i++ {
		if input[i] > greatest {
			greatest = input[i]
			index = i
		}
	}

	return index
}
