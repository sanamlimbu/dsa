package sort

import "github.com/sanamlimbu/dsa"

// Insertion sort is a simple sorting algorithm that works by iteratively
// inserting each element of an unsorted list into its correct position
// in a sorted portion of the list. It is like sorting playing cards in your hands.
//
// It is an in-place algorithm.
//
// Time complexity: O(n^2)
func InsertionSort[T dsa.Ordered](input []T, sortType dsa.SortType) {
	if sortType == dsa.DescendingSortType {
		sortDescending(input)
	} else {
		sortAscending(input)
	}
}

func sortAscending[T dsa.Ordered](input []T) {
	for i := 1; i < len(input); i++ {
		key := input[i]
		j := i - 1

		for j >= 0 && input[j] > key {
			input[j+1] = input[j]
			j--
		}

		input[j+1] = key
	}
}

func sortDescending[T dsa.Ordered](input []T) {
	for i := 1; i < len(input); i++ {
		key := input[i]
		j := i - 1

		for j >= 0 && input[j] < key {
			input[j+1] = input[j]
			j--
		}

		input[j+1] = key
	}
}
