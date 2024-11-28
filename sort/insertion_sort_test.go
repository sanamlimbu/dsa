package sort

import (
	"reflect"
	"testing"

	"github.com/sanamlimbu/dsa"
)

func TestInsertionSort(t *testing.T) {
	type args struct {
		input    []int
		sortType dsa.SortType
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "Sort ascending with integers",
			args: args{
				input:    []int{5, 2, 9, 1, 5, 6},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{1, 2, 5, 5, 6, 9},
		},
		{
			name: "Sort descending with integers",
			args: args{
				input:    []int{5, 2, 9, 1, 5, 6},
				sortType: dsa.DescendingSortType,
			},
			expected: []int{9, 6, 5, 5, 2, 1},
		},
		{
			name: "Sort ascending with empty slice",
			args: args{
				input:    []int{},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{},
		},
		{
			name: "Sort descending with empty slice",
			args: args{
				input:    []int{},
				sortType: dsa.DescendingSortType,
			},
			expected: []int{},
		},
		{
			name: "Sort ascending with one element",
			args: args{
				input:    []int{42},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{42},
		},
		{
			name: "Sort descending with one element",
			args: args{
				input:    []int{42},
				sortType: dsa.DescendingSortType,
			},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InsertionSort(tt.args.input, tt.args.sortType)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("InsertionSort(%v) = %v, want %v", tt.args.input, result, tt.expected)
			}
		})
	}
}
