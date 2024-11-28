package sort

import (
	"reflect"
	"testing"
)

func TestMergeSortAscending(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Sorted slice",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted slice",
			input:    []int{4, 1, 5, 3, 2},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Array with duplicates",
			input:    []int{3, 1, 2, 3, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element slice",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input, AscendingSortType)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("given: %v; got: %v; wanted: %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMergeSortDescending(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Sorted slice",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Unsorted slice",
			input:    []int{4, 1, 5, 3, 2},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Array with duplicates",
			input:    []int{3, 1, 2, 3, 1},
			expected: []int{3, 3, 2, 1, 1},
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element slice",
			input:    []int{42},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input, DescendingSortType)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("given: %v; got: %v; wanted: %v", tt.input, result, tt.expected)
			}
		})
	}
}
