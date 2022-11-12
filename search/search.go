package search

// Binary Search Algorthim - returns index of target, if not found returns -1
// Complexity - O(log n)
func BinaryInt(list []int, target int) int {
	if len(list) == 0 {
		return -1
	}

	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]
		if guess == target {
			return mid
		}

		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
