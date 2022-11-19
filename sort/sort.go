package sort

// SelectionInt - Selection sort algorithm
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

// QuickSort - Quick sort algorithm
// Complexity - Worst case O(n^2) , Average case O(n log n)
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := []int{}
	greater := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}
	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
