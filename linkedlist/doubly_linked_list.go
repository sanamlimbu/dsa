package linkedlist

type DoublyLinkedList[K comparable, V any] struct {
	head *doublyNode[K, V]
}

type doublyNode[K comparable, V any] struct {
	key   K
	value V
	prev  *doublyNode[K, V]
	next  *doublyNode[K, V]
}

func NewDoublyLinkedList[K comparable, V any]() *DoublyLinkedList[K, V] {
	return &DoublyLinkedList[K, V]{
		head: nil,
	}
}

func (list *DoublyLinkedList[K, V]) IsEmpty() bool {
	return list.head == nil
}

func (list *DoublyLinkedList[K, V]) InsertAtFront(key K, value V) {
	node := &doublyNode[K, V]{
		key:   key,
		value: value,
	}

	if !list.IsEmpty() {
		node.next = list.head
		list.head.prev = node
	}

	list.head = node
}

func (list *DoublyLinkedList[K, V]) InsertAtRear(key K, value V) {
	node := &doublyNode[K, V]{
		key:   key,
		value: value,
	}

	if list.IsEmpty() {
		list.head = node
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	node.prev = current
	current.next = node
}

// Insert inserts key and value pair in given position.
// Position starts at 1.
func (list *DoublyLinkedList[K, V]) Insert(key K, value V, position int) {
	if position <= 0 {
		return
	}

	if position == 1 {
		list.InsertAtFront(key, value)
		return
	}

	current := list.head

	for i := 1; i < position && current != nil; i++ {
		current = current.next
	}

	// Invalid position, greater than list length.
	if current == nil {
		return
	}

	node := &doublyNode[K, V]{
		key:   key,
		value: value,
		next:  current,
		prev:  current.prev,
	}

	current.prev.next = node
	current.prev = node
}

func (list *DoublyLinkedList[K, V]) DeleteAtFront() {
	if list.IsEmpty() {
		return
	}

	// Only one node.
	if list.head.next == nil {
		list.head = nil
		return
	}

	list.head = list.head.next
	list.head.prev = nil
}

func (list *DoublyLinkedList[K, V]) DeleteAtRear() {
	if list.IsEmpty() {
		return
	}

	// Only one node.
	if list.head.next == nil {
		list.head = nil
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.prev.next = nil
}

// Delete deletes node at given position.
// Position starts from 1.
func (list *DoublyLinkedList[K, V]) Delete(position int) {
	if position == 1 {
		list.DeleteAtFront()
		return
	}

	current := list.head

	for i := 1; i < position && current != nil; i++ {
		current = current.next
	}

	// Invalid position, greater than list length.
	if current == nil {
		return
	}

	// Not last node.
	if current.next != nil {
		current.next.prev = current.prev
	}

	// Not first node.
	if current.prev != nil {
		current.prev.next = current.next
	}
}
