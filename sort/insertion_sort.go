package sort

import "github.com/sanamlimbu/dsa"

// Insertion sort is a simple sorting algorithm that works by iteratively
// inserting each element of an unsorted list into its correct position
// in a sorted portion of the list. It is like sorting playing cards in your hands.
func InsertionSort[T dsa.Ordered](input []T, sortType dsa.SortType) []T {
	result := make([]T, len(input))
	copy(result, input)

	if sortType == dsa.DescendingSortType {
		sortDescending(result)
	} else {
		sortAscending(result)
	}

	return result
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
