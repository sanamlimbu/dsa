package binarysearch

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type Numeric interface {
	Int | Float
}

type Input interface {
	Numeric | ~string
}

// Search performs a binary search on a sorted slice to find the index of the target element.
// If the target element is found, it returns its index; otherwise, it returns -1.
// The input slice must be sorted, and the element type T must support all relational operators:
// <, <=, >, >=, ==, and !=
// Time complexity: O(log n)
func Search[T Input](list []T, target T) int {
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
