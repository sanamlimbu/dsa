package sort

import (
	"reflect"
	"testing"

	"github.com/sanamlimbu/dsa"
)

func TestHeapSort(t *testing.T) {
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
				input:    []int{5, 4, 3, 2, 1, 0},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "Sort ascending with integers",
			args: args{
				input:    []int{5, 2, 9, 1, 1, 7, 5},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{1, 1, 2, 5, 5, 7, 9},
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
			name: "Sort ascending with one element",
			args: args{
				input:    []int{42},
				sortType: dsa.AscendingSortType,
			},
			expected: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.input, tt.args.sortType)
			if !reflect.DeepEqual(tt.args.input, tt.expected) {
				t.Errorf("HeapSort(%v) = %v, want %v", tt.args.input, tt.args.input, tt.expected)
			}
		})
	}
}
