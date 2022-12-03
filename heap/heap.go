package heap

import "fmt"

type MaxHeap struct {
	array []int
}

// Insrt adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() (int, error) {
	// Array is empty?
	if len(h.array) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	extracted := h.array[0]

	l := len(h.array) - 1
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted, nil
}

func (h *MaxHeap) IsEmpty() bool {
	return len(h.array) == 0
}

func (h *MaxHeap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.array[0], nil
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l

		} else { // when right child is larger
			childToCompare = r
		}

		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// Get parent index
// i is index of the child
func parent(i int) int {
	return (i - 1) / 2
}

// Get left child index
// i is index of the parent
func left(i int) int {
	return 2*i + 1
}

// Get right child index
// i is index of the parent
func right(i int) int {
	return 2*i + 2
}

// Swap keys in the array
func (h *MaxHeap) swap(i, j int) {
	temp := h.array[i]
	h.array[i] = h.array[j]
	h.array[i] = temp

	// golang specific operation
	// h.array[i], h.array[j] = h.array[j], h.array[i]
}
