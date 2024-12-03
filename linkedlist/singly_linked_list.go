package linkedlist

type SinglyLinkedList[K comparable, V any] struct {
	head *node[K, V]
}

func NewSinglyLinkedList[K comparable, V any]() *SinglyLinkedList[K, V] {
	return &SinglyLinkedList[K, V]{
		head: nil,
	}
}

type node[K comparable, V any] struct {
	key   K
	value V
	next  *node[K, V]
}

func (list *SinglyLinkedList[K, V]) IsEmpty() bool {
	return list.head == nil
}

// InsertAtPosition inserts key and value pair in the given position.
// Position starts from 1.
// Inserts at front when list is empty or position is less than or equal to 0.
// Inserts at last when position is out of range.
// Inserts at given position, when position is in the range.
func (list *SinglyLinkedList[K, V]) InsertAtPosition(key K, value V, position int) {
	node := &node[K, V]{
		key:   key,
		value: value,
		next:  nil,
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

// Insert inserts key and value pair.
// If key is found then value is overwritten.
// If not found then new node is added at the end.
func (list *SinglyLinkedList[K, V]) Insert(key K, value V) {
	node := &node[K, V]{
		key:   key,
		value: value,
		next:  nil,
	}

	if list.IsEmpty() {
		list.head = node
		return
	}

	current := list.head

	if current.key == key {
		current.value = value
		return
	}

	for current.next != nil {
		if current.next.key == key {
			current.next.value = value
			return
		}

		current = current.next
	}

	current.next = node
}

// Delete deletes node in the given position.
// Position starts from 1.
func (list *SinglyLinkedList[K, V]) Delete(position int) {
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

// DeleteKey deletes node with given key.
func (list *SinglyLinkedList[K, V]) DeleteKey(key K) {
	if list.IsEmpty() {
		return
	}

	current := list.head

	// Head node matches.
	if current.key == key {
		list.head = list.head.next
		return
	}

	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			return
		}

		current = current.next
	}
}

// Length returns number of nodes in the linked list.
func (list *SinglyLinkedList[K, V]) Length() int {
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

type pair[K comparable, V any] struct {
	key   K
	value V
}

// Slice returns slice of key-value pairs in the linked list.
func (list *SinglyLinkedList[K, V]) Slice() []pair[K, V] {
	if list.IsEmpty() {
		return nil
	}

	current := list.head
	result := []pair[K, V]{{
		key:   current.key,
		value: current.value,
	}}

	for current.next != nil {
		current = current.next
		result = append(result, pair[K, V]{
			key:   current.key,
			value: current.value,
		})
	}

	return result
}

// Search searches given key on linked list.
// Returns true when found and false when not found.
// Returns value as well.
func (list *SinglyLinkedList[K, V]) Search(key K) (V, bool) {
	current := list.head

	var value V

	for current != nil {
		if current.key == key {
			return current.value, true
		}

		current = current.next
	}

	return value, false
}
