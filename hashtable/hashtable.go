package hashtable

import (
	"github.com/sanamlimbu/dsa/linkedlist"
)

const arraySize int = 10

type HashTable[V any] struct {
	array [arraySize]*bucket[V]
}

func NewHashTable[V any]() *HashTable[V] {
	ht := &HashTable[V]{}

	for i := range ht.array {
		ht.array[i] = &bucket[V]{}
	}

	return ht
}

// Bucket is a singly linked list.
type bucket[V any] struct {
	linkedlist.SinglyLinkedList[string, V]
}

// hash takes key and returns int code.
// Returned int code is the range 0 <= code < arraySize.
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % arraySize
}

// Set takes in key and value, and inserts into hash table.
// If key is already there it overwrites the value.
func (hashTable *HashTable[V]) Set(key string, value V) {
	index := hash(key)
	hashTable.array[index].Insert(key, value)
}

// Get will take a key and returns value and true if found in hash table.
func (hashTable *HashTable[V]) Get(key string) (V, bool) {
	index := hash(key)
	return hashTable.array[index].Search(key)
}

// Delete will take a key and deletes from hash table.
func (hashTable *HashTable[V]) Delete(key string) {
	index := hash(key)
	hashTable.array[index].DeleteKey(key)
}
