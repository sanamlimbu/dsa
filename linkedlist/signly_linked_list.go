package linkedlist

type SinglyLinkedList[T comparable] struct {
	head *node[T]
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
	}
}

type node[T any] struct {
	data T
	next *node[T]
}

func (list *SinglyLinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

// Insert inserts data in the given position.
// Position starts from 1.
// Inserts at front when list is empty or position is less than or equal to 0.
// Inserts at last when position is out of range.
// Inserts at given position, when position is in the range.
func (list *SinglyLinkedList[T]) Insert(data T, position int) {
	node := &node[T]{
		data: data,
		next: nil,
	}

	// Invalid position or empty list.
	// Insert at front.
	if position <= 0 || list.IsEmpty() {
		list.head = node
		return
	}

	// Insertion at front.
	// Make new node the head node.
	if position == 1 {
		node.next = list.head
		list.head = node
		return
	}

	// Insert at the given position.
	// If position is out of range, insert at last.
	current := list.head

	for i := 1; i < position-1 && current.next != nil; i++ {
		current = current.next
	}

	node.next = current.next
	current.next = node
}

// Delete deletes node in the given position.
// Position starts from 1.
func (list *SinglyLinkedList[T]) Delete(position int) {
	if position <= 0 || list.IsEmpty() {
		return
	}

	if position == 1 {
		list.head = list.head.next
		return
	}

	current := list.head

	for i := 1; i < position-1 && current.next != nil; i++ {
		current = current.next
	}

	current.next = current.next.next
}

// Length returns number of nodes in the linked list.
func (list *SinglyLinkedList[T]) Length() int {
	if list.IsEmpty() {
		return 0
	}

	count := 1

	current := list.head

	for current.next != nil {
		current = current.next
		count++
	}

	return count
}

// Slice returns slice of data in the linked list.
func (list *SinglyLinkedList[T]) Slice() []T {
	if list.IsEmpty() {
		return nil
	}

	current := list.head
	result := []T{current.data}

	for current.next != nil {
		current = current.next
		result = append(result, current.data)
	}

	return result
}

// Search searches given element on linked list.
// Returns true when found and false when not found.
func (list *SinglyLinkedList[T]) Search(data T) bool {
	current := list.head

	for current != nil {
		if current.data == data {
			return true
		}

		current = current.next
	}

	return false
}
