package sort

/*
* Merge sort algorithm is based on divide and conquer paradigm
* time complexity is O(n log n)
 */
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	midIndex := len(arr) / 2
	leftHalf := arr[:midIndex]
	rightHalf := arr[midIndex:]

	leftHalf = MergeSort(leftHalf)
	rightHalf = MergeSort(rightHalf)

	return merge(leftHalf, rightHalf)
}

func merge(leftHalf, rightHalf []int) []int {
	result := []int{}
	leftSize := len(leftHalf)
	rightSize := len(rightHalf)

	i, j := 0, 0
	for i < leftSize && j < rightSize {
		if leftHalf[i] <= rightHalf[j] {
			result = append(result, leftHalf[i])
			i++
		} else {
			result = append(result, rightHalf[j])
			j++
		}
	}

	for ; i < leftSize; i++ {
		result = append(result, leftHalf[i])
	}

	for ; j < rightSize; j++ {
		result = append(result, rightHalf[j])
	}

	return result
}
