package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{9, 1, 7, 5, 3, 6}
	arrWanted := []int{1, 3, 5, 6, 7, 9}
	arrGot := MergeSort(arr)
	compare(arrWanted, arrGot)

	arr = []int{9, 8, 6, 3, 1}
	arrWanted = []int{1, 3, 6, 8, 9}
	arrGot = MergeSort(arr)
	compare(arrWanted, arrGot)

	arr = []int{9, 1, 6, 7, 2}
	arrWanted = []int{1, 2, 6, 7, 9}
	arrGot = MergeSort(arr)
	compare(arrWanted, arrGot)
}
