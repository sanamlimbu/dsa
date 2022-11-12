package sort

// SlectionInt - Selection sort algorithm
// Complexity - O(n^2)
func SelectionInt(arr []int) []int {
	newArr := make([]int, 0, len(arr))
	for range arr {
		smallest := FindSmallest(arr)
		newArr = append(newArr, arr[smallest])

		// delete smallest element
		arr[smallest] = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
	}
	return newArr
}

// FindSmallest returns index of smallest element
func FindSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}
