package recursion

import "math"

// SumInt - uses recursion to find sum of array
func SumInt(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + SumInt(arr[1:])
}

// CountInt - uses recursion to find number of items in array
func CountInt(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + CountInt(arr[1:])
}

// MaximumInt - uses recursion to find maximum number
func MaximumInt(arr []int, n int) int {
	if n == 1 {
		return arr[0]
	}
	lastItem := float64(arr[n-1])
	return int(math.Max(lastItem, float64(MaximumInt(arr, n-1))))
}

// BinarySearch - resurisve binary search
func BinarySearch(arr []int, target int, low int, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2

	if target == arr[mid] {
		return mid
	} else if target > arr[mid] {
		return BinarySearch(arr, target, mid+1, high)

	} else {
		return BinarySearch(arr, target, low, mid-1)
	}
}
