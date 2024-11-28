package sort

import (
	"reflect"
	"testing"

	"github.com/sanamlimbu/dsa"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		sortType dsa.SortType
		expected []int
	}{
		{
			name:     "Ascending order",
			input:    []int{3, 1, 4, 1, 5, 9, 2},
			sortType: dsa.AscendingSortType,
			expected: []int{1, 1, 2, 3, 4, 5, 9},
		},
		{
			name:     "Descending order",
			input:    []int{3, 1, 4, 1, 5, 9, 2},
			sortType: dsa.DescendingSortType,
			expected: []int{9, 5, 4, 3, 2, 1, 1},
		},
		{
			name:     "Empty input",
			input:    []int{},
			sortType: dsa.AscendingSortType,
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			sortType: dsa.AscendingSortType,
			expected: []int{42},
		},
		{
			name:     "Already sorted ascending",
			input:    []int{1, 2, 3, 4, 5},
			sortType: dsa.AscendingSortType,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Already sorted descending",
			input:    []int{5, 4, 3, 2, 1},
			sortType: dsa.DescendingSortType,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "All elements the same",
			input:    []int{7, 7, 7, 7},
			sortType: dsa.AscendingSortType,
			expected: []int{7, 7, 7, 7},
		},
		{
			name:     "Negative numbers",
			input:    []int{-3, -1, -4, -1, -5, -9, -2},
			sortType: dsa.AscendingSortType,
			expected: []int{-9, -5, -4, -3, -2, -1, -1},
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{-3, 1, -4, 1, -5, 9, 2},
			sortType: dsa.DescendingSortType,
			expected: []int{9, 2, 1, 1, -3, -4, -5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := QuickSort(tt.input, tt.sortType)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("QuickSort(%v, %v) = %v; want %v", tt.input, tt.sortType, result, tt.expected)
			}
		})
	}
}
