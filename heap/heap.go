package heap

import (
	"errors"

	"github.com/sanamlimbu/dsa"
)

type Heap[T dsa.Ordered] interface {
	// Insert inserts an element to the heap maintaining heap property.
	Insert(T)

	// Extract returns root of the heap and detach it.
	Extract() (T, error)

	// IsEmpty checks if heap is empty.
	IsEmpty() bool

	// Peek returns root of the heap.
	Peek() (T, error)

	// Slice returns copy of underlying slice.
	Slice() []T

	heapifyUp(int)
	heapifyDown(int)
	swap(int, int)
}

type HeapType string

const (
	MaxHeapType HeapType = "max-heap"
	MinHeapType HeapType = "min-heap"
)

var ErrHeapEmpty = errors.New("heap is empty")

func NewHeap[T dsa.Ordered](t HeapType, elements []T) Heap[T] {
	if t == MaxHeapType {
		return newMaxHeap(elements)
	}

	return newMaxHeap(elements)
}

type maxHeap[T dsa.Ordered] struct {
	slice []T
}

func newMaxHeap[T dsa.Ordered](elements []T) *maxHeap[T] {
	heap := &maxHeap[T]{
		slice: make([]T, 0),
	}

	for _, element := range elements {
		heap.Insert(element)
	}

	return heap
}

// Slice returns copy of underlying slice.
func (h *maxHeap[T]) Slice() []T {
	slice := make([]T, len(h.slice))

	copy(slice, h.slice)

	return slice
}

// Insert adds an element to the heap maintaining the heap property.
func (h *maxHeap[T]) Insert(element T) {
	h.slice = append(h.slice, element)
	h.heapifyUp(len(h.slice) - 1)
}

// Extract returns root element, and removes it from the heap.
func (h *maxHeap[T]) Extract() (T, error) {
	var element T

	if h.IsEmpty() {
		return element, ErrHeapEmpty
	}

	extracted := h.slice[0]

	lastIndex := len(h.slice) - 1

	h.slice[0] = h.slice[lastIndex] // move last element to root

	h.slice = h.slice[:lastIndex] // delete last element

	h.heapifyDown(0)

	return extracted, nil
}

func (h *maxHeap[T]) IsEmpty() bool {
	return len(h.slice) == 0
}

// Peek returns root element in the heap.
func (h *maxHeap[T]) Peek() (T, error) {
	var element T

	if h.IsEmpty() {
		return element, ErrHeapEmpty
	}

	return h.slice[0], nil
}

// heapifyDown will heapify top to bottom, starting from given index 'i'.
func (h *maxHeap[T]) heapifyDown(i int) {
	last := len(h.slice) - 1

	l, r := left(i), right(i)

	var childToCompare int

	// loop while index has at least one child
	for l <= last {
		switch {
		case l == last: // when left child is the only child
			childToCompare = l

		case h.slice[l] > h.slice[r]: // when left child is greater than right child
			childToCompare = l

		case h.slice[r] > h.slice[l]: // when right child is greater than left child
			childToCompare = r
		}

		// when current element is not smaller than it's child to compare
		if h.slice[i] >= h.slice[childToCompare] {
			return
		}

		// when current element is smaller than it's child to compare, swap
		h.swap(i, childToCompare)
		i = childToCompare
		l, r = left(i), right(i)
	}
}

// heapifyUp will heapify from bottom to top, starting from given 'i' index.
func (h *maxHeap[T]) heapifyUp(i int) {
	j := parent(i)

	for h.slice[j] < h.slice[i] {
		h.swap(j, i)
		i = j
	}
}

// swap swaps element in index 'i' and element in index 'j'.
func (h *maxHeap[T]) swap(i, j int) {
	h.slice[i], h.slice[j] = h.slice[j], h.slice[i]
}

// parent returns index of the parent for given 'i' child index.
func parent(i int) int {
	return (i - 1) / 2
}

// left returns the index of left child for given parent index 'i'
func left(i int) int {
	return 2*i + 1
}

// right returns the index of right child for given parent index 'i'
func right(i int) int {
	return 2*i + 2
}
