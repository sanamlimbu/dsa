package heap

import (
	"reflect"
	"testing"
)

func TestBuildMaxHeap(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Already a max heap",
			input:    []int{20, 10, 15},
			expected: []int{20, 10, 15},
		},
		{
			name:     "Unsorted slice",
			input:    []int{10, 20, 5, 6, 1},
			expected: []int{20, 10, 5, 6, 1},
		},
		{
			name:     "Negative numbers",
			input:    []int{-10, -20, -5, -1, -6},
			expected: []int{-1, -5, -10, -20, -6},
		},
		{
			name:     "Mixed positive and negative numbers",
			input:    []int{3, -2, 15, 7, -8},
			expected: []int{15, 7, 3, -2, -8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewHeap(MaxHeapType, tt.input)
			if !reflect.DeepEqual(result.Slice(), tt.expected) {
				t.Errorf("For input %v, expected %v, got %v", tt.input, tt.expected, result)
			}
		})
	}
}
