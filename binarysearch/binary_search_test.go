package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{arr: []int{-9, 0, 1, 5, 10, 12}, target: 5, want: 3},
		{arr: []int{0, 2, 4, 9, 23, 70}, target: 23, want: 4},
		{arr: []int{0, 2, 3, 10, 29, 31}, target: 99, want: -1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Searching %d in list", tc.target), func(t *testing.T) {
			got := Search(tc.arr, tc.target)

			if got != tc.want {
				t.Errorf("wanted index %d but got %d", tc.want, got)
			}
		})
	}
}

func TestSearchByRecursion(t *testing.T) {
	testCases := []struct {
		arr    []int
		target int
		want   int
	}{
		{arr: []int{-9, 0, 1, 5, 10, 12}, target: 5, want: 3},
		{arr: []int{0, 2, 4, 9, 23, 70}, target: 23, want: 4},
		{arr: []int{0, 2, 3, 10, 29, 31}, target: 99, want: -1},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(fmt.Sprintf("Searching %d in list", tc.target), func(t *testing.T) {
			got := SearchByRecursion(tc.arr, tc.target, 0, len(tc.arr)-1)

			if got != tc.want {
				t.Errorf("wanted index %d but got %d", tc.want, got)
			}
		})
	}
}
