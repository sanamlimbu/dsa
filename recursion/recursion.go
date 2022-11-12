package recursion

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
