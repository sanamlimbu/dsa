package queue

import (
	"errors"
)

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() (T, error)
	Peek() (T, error)
	IsEmpty() bool
}

type QueueImplementation string

const (
	SliceQueueImplementation      QueueImplementation = "slice"
	LinkedListQueueImplementation QueueImplementation = "linked-list"
)

var ErrQueueEmpty = errors.New("queue is empty")

// NewQueue creates an empty queue.
func NewQueue[T any](implementation QueueImplementation) Queue[T] {
	if implementation == SliceQueueImplementation {
		return newSliceQueue[T]()
	}

	return newLinkedListQueue[T]()
}

// sliceQueue represents a queue implemented using a slice.
type sliceQueque[T any] struct {
	slice []T
}

// newSliceQueue creates an empty queue i.e underlying slice length is zero.
func newSliceQueue[T any]() *sliceQueque[T] {
	return &sliceQueque[T]{
		slice: make([]T, 0),
	}
}

func (q *sliceQueque[T]) Enqueue(element T) {
	q.slice = append(q.slice, element)
}

func (q *sliceQueque[T]) Dequeue() (T, error) {
	var element T

	length := len(q.slice)

	if length == 0 {
		return element, ErrQueueEmpty
	}

	element = q.slice[0]

	slice := make([]T, length-1)

	copy(slice, q.slice[1:])

	q.slice = slice

	return element, nil
}

func (q *sliceQueque[T]) Peek() (T, error) {
	var element T

	if q.IsEmpty() {
		return element, ErrQueueEmpty
	}

	return q.slice[0], nil
}

func (q *sliceQueque[T]) IsEmpty() bool {
	return len(q.slice) == 0
}

// linkedListQueue is queue implementation with singly linked list.
type linkedListQueue[T any] struct {
	front *node[T]
	rear  *node[T]
}

type node[T any] struct {
	element T
	next    *node[T]
}

// newLinkedListQueue creates and initializes an empty queue.
func newLinkedListQueue[T any]() *linkedListQueue[T] {
	return &linkedListQueue[T]{
		front: nil,
		rear:  nil,
	}
}

func (q *linkedListQueue[T]) Enqueue(element T) {
	node := &node[T]{element: element}

	if q.IsEmpty() {
		q.front = node
		q.rear = node
		return
	}

	q.rear.next = node
	q.rear = node
}

func (q *linkedListQueue[T]) Dequeue() (T, error) {
	var element T

	if q.IsEmpty() {
		return element, ErrQueueEmpty
	}

	frontNode := q.front
	element = frontNode.element

	q.front = frontNode.next
	frontNode.next = nil

	if q.front == nil {
		q.rear = nil
	}

	return element, nil
}

func (q *linkedListQueue[T]) Peek() (T, error) {
	var element T

	if q.IsEmpty() {
		return element, ErrQueueEmpty
	}

	element = q.front.element

	return element, nil
}

func (q *linkedListQueue[T]) IsEmpty() bool {
	return q.front == nil && q.rear == nil
}
