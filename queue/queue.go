package queue

import "fmt"

const ArraySize = 5

type QuequeArray struct {
	front    int
	rear     int
	capacity int
	array    [ArraySize]int
}

func NewQueueArray() QuequeArray {
	q := QuequeArray{front: 0, rear: -1, capacity: ArraySize}
	return q
}

func (q *QuequeArray) Enqueue(value int) error {
	if (q.rear + 1) < q.capacity {
		q.rear++
		q.array[q.rear] = value
		return nil
	}
	return fmt.Errorf("queue is full")
}

func (q *QuequeArray) Dequeue() (int, error) {
	if q.rear < 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	value := q.array[q.front]

	// move remaning elememts on index down
	for i := 0; i < q.rear; i++ {
		q.array[i] = q.array[i+1]
	}
	q.rear--

	return value, nil
}

func (q *QuequeArray) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.array[q.front], nil
}

func (q *QuequeArray) IsEmpty() bool {
	if q.rear < 0 {
		return true
	}
	return false
}

func (q *QuequeArray) IsFull() bool {
	if (q.rear + 1) < q.capacity {
		return false
	}
	return true
}

// QueueLinkedList is queue implementation with linked list
type QuequeLinkedList struct {
	front *QueueNode
	head  *QueueNode // head is also rear and points last node
}

type QueueNode struct {
	value int
	next  *QueueNode
}

func (q *QuequeLinkedList) Enqueue(value int) {
	newNode := QueueNode{value: value}
	// Empty queue ?
	if q.head == nil {
		q.front = &newNode
	}
	newNode.next = q.head
	q.head = &newNode
}

func (q *QuequeLinkedList) Dequeue() (int, error) {
	if q.head == nil {
		return 0, fmt.Errorf("queue is empty")
	}

	// Get value from front of queue
	value := q.front.value

	// Only one node ?
	if q.head == q.front {
		q.head = nil
	}

	// Detach the front and make new front node of the queue
	previousNode := q.head
	for previousNode != nil {
		if previousNode.next == q.front {
			previousNode.next = nil
			break
		}
		previousNode = previousNode.next
	}
	q.front = previousNode
	return value, nil
}

func (q *QuequeLinkedList) Peek() (int, error) {
	if q.head == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.front.value, nil
}

func (q *QuequeLinkedList) IsEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}
