package binarysearch

import "github.com/sanamlimbu/dsa"

// Search performs a binary search on a sorted slice to find the index of the target element using iteration.
// If the target element is found, it returns its index; otherwise, it returns -1.
// The input slice must be sorted, and the element type T must support all relational operators:
// <, <=, >, >=, ==, and !=
// Time complexity: O(log n)
func Search[T dsa.Ordered](list []T, target T) int {
	if len(list) == 0 {
		return -1
	}

	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		switch {
		case guess == target:
			return mid
		case guess > target:
			high = mid - 1
		case guess < target:
			low = mid + 1
		}
	}

	return -1
}

// SearchByRecursion performs a binary search on a sorted slice to find the index of the target element using recursion.
// If the target element is found, it returns its index; otherwise, it returns -1.
// The input slice must be sorted, and the element type T must support all relational operators:
// <, <=, >, >=, ==, and !=
// Time complexity: O(log n)
func SearchByRecursion[T dsa.Ordered](list []T, target T, low, high int) int {
	if len(list) == 0 || low > high {
		return -1
	}

	mid := low + (high-low)/2

	switch {
	case target == list[mid]:
		return mid
	case target < list[mid]:
		return SearchByRecursion(list, target, low, mid-1)
	case target > list[mid]:
		return SearchByRecursion(list, target, mid+1, high)
	default:
		return -1
	}
}
