package heap

import (
	"errors"

	"github.com/sanamlimbu/dsa"
)

// A Heap is a complete binary tree data structure.
// A complete binary tree is a special type of binary tree where all the levels of the tree are filled completely
// except the lowest level nodes which are filled from as left as possible.
//
// It satisfies the heap property: for every node, the value of its children is greater than or equal to its own value.
// It is often used to implement priority queues because it allows efficient retrieval of the highest (or lowest) priority element.
//
// The heap property ensures that:
// In a max-heap, the value of each parent node is greater than or equal to the values of its children, and the largest value is at the root of the tree.
// In a min-heap, the value of each parent node is smaller than or equal to the values of its children, and the smallest value is at the root of the tree.

// 		Min-Heap         		Max-Heap
//
//	       1                       11
//	     /   \                   /    \
//	    2     3                 10      9
//	   / \   / \               /  \    /  \
//	  4   5 6   7            8     7  6    5
//	 / \ / \                / \  / \
//	8  9 10 11            4  3 2   1

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

	// loop until bottom of the tree is reached
	// don't overflow
	for l <= last {
		switch {
		case l == last: // when left child is the only child
			childToCompare = l

		case h.slice[l] > h.slice[r]: // when left child is greater than right child
			childToCompare = l

		case h.slice[r] > h.slice[l]: // when right child is greater than left child
			childToCompare = r

		case h.slice[l] == h.slice[r]: // when left child is equal to right right
			childToCompare = l
		}

		// when current element is greater or equal to it's child to compare
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
	for i > 0 {
		j := parent(i)

		if h.slice[j] >= h.slice[i] {
			return
		}

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
