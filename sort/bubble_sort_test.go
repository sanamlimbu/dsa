package sort

import (
	"reflect"
	"testing"

	"github.com/sanamlimbu/dsa"
)

func TestBubbleSort(t *testing.T) {
	type args[T dsa.Ordered] struct {
		input    []T
		sortType dsa.SortType
	}

	tests := []struct {
		name     string
		args     args[int]
		expected []int
	}{
		{
			name: "Ascending order",
			args: args[int]{
				input:    []int{5, 3, 8, 6, 2},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{2, 3, 5, 6, 8},
		},
		{
			name: "Descending order",
			args: args[int]{
				input:    []int{5, 3, 8, 6, 2},
				sortType: dsa.DescendingSortType,
			},
			expected: []int{8, 6, 5, 3, 2},
		},
		{
			name: "Empty slice",
			args: args[int]{
				input:    []int{},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{},
		},
		{
			name: "Single element",
			args: args[int]{
				input:    []int{42},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{42},
		},
		{
			name: "Already sorted ascending",
			args: args[int]{
				input:    []int{1, 2, 3, 4, 5},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Already sorted descending",
			args: args[int]{
				input:    []int{5, 4, 3, 2, 1},
				sortType: dsa.DescendingSortType,
			},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name: "All elements identical",
			args: args[int]{
				input:    []int{7, 7, 7, 7},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{7, 7, 7, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.args.input))
			copy(input, tt.args.input)

			BubbleSort(tt.args.input, tt.args.sortType)

			if !reflect.DeepEqual(tt.args.input, tt.expected) {
				t.Errorf("BubbleSort(%v, %v) = %v, want %v",
					tt.args.input, tt.args.sortType, tt.args.input, tt.expected)
			}
		})
	}
}
