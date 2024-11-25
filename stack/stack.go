package stack

import "fmt"

const StackSize = 10

// SliceStack represents a stack implemented using a slice.
// The underlying slice dynamically grows as needed to accommodate new elements.
type SliceStack[T any] struct {
	slice []T
}

// NewSliceStack creates an empty stack.
// The underlying slice length is zero.
func NewSliceStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{
		slice: make([]T, 0),
	}
}

func (s *SliceStack[T]) Push(element T) {
	s.slice = append(s.slice, element)
}

func (s *SliceStack[T]) Pop() (T, error) {
	var element T

	length := len(s.slice)

	if length == 0 {
		return element, fmt.Errorf("stack is empty")
	}

	element = s.slice[length-1]

	s.slice = s.slice[:length-1]

	return element, nil
}

func (s *SliceStack[T]) Peek() (T, error) {
	var element T

	length := len(s.slice)

	if length == 0 {
		return element, fmt.Errorf("stack is empty")
	}

	element = s.slice[length-1]

	return element, nil
}

func (s *SliceStack[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

// LinkedListStack represents a stack implemented using a singly linked list.
type LinkedListStack[T any] struct {
	head *Node[T]
}

// Node is a singly linked list node.
type Node[T any] struct {
	element T
	next    *Node[T]
}

// NewLinkedListStack creates an empty stack.
// The underlying singly linked list's head points to nil.
func NewLinkedListStack[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{
		head: nil,
	}
}

func (s *LinkedListStack[T]) Push(element T) {
	node := &Node[T]{element: element}

	node.next = s.head

	s.head = node
}

func (s *LinkedListStack[T]) Pop() (T, error) {
	var element T

	if s.head == nil {
		return element, fmt.Errorf("stack is empty")
	}

	element = s.head.element
	s.head = s.head.next

	return element, nil
}

func (s *LinkedListStack[T]) Peek() (T, error) {
	var element T

	if s.head == nil {
		return element, fmt.Errorf("stack is empty")
	}

	return s.head.element, nil
}

func (s *LinkedListStack[T]) IsEmpty() bool {
	return s.head == nil
}
