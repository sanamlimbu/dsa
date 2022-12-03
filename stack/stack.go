package stack

import "fmt"

const StackSize = 10

// StackArray is stack implemention with array
type StackArray struct {
	top   int
	max   int
	array []int
}

func NewStackArray() StackArray {
	stack := StackArray{}
	stack.top = -1
	stack.max = StackSize
	array := [StackSize]int{}
	stack.array = array[:]
	return stack
}

func (s *StackArray) Push(value int) error {
	if (s.top + 1) < s.max {
		s.top++
		s.array[s.top] = value
		return nil
	}
	return fmt.Errorf("stack is full")
}

func (s *StackArray) Pop() (int, error) {
	if s.top < 0 {
		return 0, fmt.Errorf("stack is empty")

	}
	value := s.array[s.top]
	s.top--
	return value, nil
}

func (s *StackArray) Peek() (int, error) {
	if s.top < 0 {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.array[s.top], nil
}

func (s *StackArray) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *StackArray) IsFull() bool {
	if (s.top + 1) == s.max {
		return true
	}
	return false
}

// StackLinkedList is stack implementation with linked list
type StackLinkedList struct {
	head *StackNode
}

// StackNode is a singly linked list node
type StackNode struct {
	value int
	next  *StackNode
}

func (s *StackLinkedList) Push(value int) {
	newNode := StackNode{value: value}
	newNode.next = s.head
	s.head = &newNode
}

func (s *StackLinkedList) Pop() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("stack is empty")
	}
	s.head = s.head.next
	s.head.next = nil
	return s.head.value, nil
}

func (s *StackLinkedList) Peek() (int, error) {
	if s.head == nil {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.head.value, nil
}

func (s *StackLinkedList) IsEmpty() bool {
	if s.head == nil {
		return true
	}
	return false
}
