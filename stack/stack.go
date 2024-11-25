package stack

import "fmt"

type Stack[T any] interface {
	Push(T)
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
}

type StackImplementation string

const (
	SliceStackImplementation      StackImplementation = "stack"
	LinkedListStackImplementation StackImplementation = "linked-list"
)

// NewStack creates an empty stack.
func NewStack[T any](implementation StackImplementation) Stack[T] {
	if implementation == SliceStackImplementation {
		return newSliceStack[T]()
	}

	return newLinkedListStack[T]()
}

// sliceStack represents a stack implemented using a slice.
// The underlying slice dynamically grows as needed to accommodate new elements.
type sliceStack[T any] struct {
	slice []T
}

// newSliceStack creates an empty stack.
// The underlying slice length is zero.
func newSliceStack[T any]() *sliceStack[T] {
	return &sliceStack[T]{
		slice: make([]T, 0),
	}
}

func (s *sliceStack[T]) Push(element T) {
	s.slice = append(s.slice, element)
}

func (s *sliceStack[T]) Pop() (T, error) {
	var element T

	length := len(s.slice)

	if length == 0 {
		return element, fmt.Errorf("stack is empty")
	}

	element = s.slice[length-1]

	s.slice = s.slice[:length-1]

	return element, nil
}

func (s *sliceStack[T]) Peek() (T, error) {
	var element T

	length := len(s.slice)

	if length == 0 {
		return element, fmt.Errorf("stack is empty")
	}

	element = s.slice[length-1]

	return element, nil
}

func (s *sliceStack[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

// linkedListStack represents a stack implemented using a singly linked list.
type linkedListStack[T any] struct {
	head *node[T]
}

// node is a singly linked list node.
type node[T any] struct {
	element T
	next    *node[T]
}

// newLinkedListStack creates an empty stack.
// The underlying singly linked list's head points to nil.
func newLinkedListStack[T any]() *linkedListStack[T] {
	return &linkedListStack[T]{
		head: nil,
	}
}

func (s *linkedListStack[T]) Push(element T) {
	node := &node[T]{element: element}

	node.next = s.head

	s.head = node
}

func (s *linkedListStack[T]) Pop() (T, error) {
	var element T

	if s.head == nil {
		return element, fmt.Errorf("stack is empty")
	}

	element = s.head.element
	s.head = s.head.next

	return element, nil
}

func (s *linkedListStack[T]) Peek() (T, error) {
	var element T

	if s.head == nil {
		return element, fmt.Errorf("stack is empty")
	}

	return s.head.element, nil
}

func (s *linkedListStack[T]) IsEmpty() bool {
	return s.head == nil
}
