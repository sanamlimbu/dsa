package sort

import "github.com/sanamlimbu/dsa"

type SortType string

const (
	AscendingSortType  SortType = "ascending"
	DescendingSortType SortType = "descending"
)

// Merge sort is a sorting algorithm that follows the divide-and-conquer approach.
// It works by recursively dividing the input array into smaller subarrays
// and sorting those subarrays then merging them back together to obtain the sorted array.
// Time complexity: O(nlogn)
func MergeSort[T dsa.Ordered](input []T, sortType SortType) []T {
	if len(input) < 2 {
		return input
	}

	// copying so input is not affected
	arr := make([]T, len(input))
	copy(arr, input)

	midIndex := len(arr) / 2
	leftHalf := arr[:midIndex]
	rightHalf := arr[midIndex:]

	// left half comes out sorted
	leftHalf = MergeSort(leftHalf, sortType)

	// right half comes out sorted
	rightHalf = MergeSort(rightHalf, sortType)

	// combined value in descending order
	if sortType == DescendingSortType {
		return mergeDescending(leftHalf, rightHalf)
	}

	// combined value in ascending order
	return mergeAscending(leftHalf, rightHalf)
}

// leftHalf and rightHalf are already in ascending order
func mergeAscending[T dsa.Ordered](leftHalf, rightHalf []T) []T {
	result := make([]T, 0)
	leftSize := len(leftHalf)
	rightSize := len(rightHalf)

	var i, j int

	for i < leftSize && j < rightSize {
		switch {
		case leftHalf[i] <= rightHalf[j]:
			result = append(result, leftHalf[i])
			i++
		default:
			result = append(result, rightHalf[j])
			j++
		}
	}

	// merge remaining values from left half
	for ; i < leftSize; i++ {
		result = append(result, leftHalf[i])
	}

	// merge remaining values from right half
	for ; j < rightSize; j++ {
		result = append(result, rightHalf[j])
	}

	return result
}

// leftHalf and rightHalf are already in descending order
func mergeDescending[T dsa.Ordered](leftHalf, rightHalf []T) []T {
	result := make([]T, 0)
	leftSize := len(leftHalf)
	rightSize := len(rightHalf)

	var i, j int

	for i < leftSize && j < rightSize {
		switch {
		case leftHalf[i] < rightHalf[j]:
			result = append(result, rightHalf[j])
			j++
		default:
			result = append(result, leftHalf[i])
			i++
		}
	}

	// merge remaining values from left half
	for ; i < leftSize; i++ {
		result = append(result, leftHalf[i])
	}

	// merge remaining values from right half
	for ; j < rightSize; j++ {
		result = append(result, rightHalf[j])
	}

	return result
}
