package hashtable

import "fmt"

const ArraySize int = 5

// HashTable has array of size 5 to store key value pairs
type HashTable struct {
	array [ArraySize]*Bucket
}

// Bucket is a singly linked list
type Bucket struct {
	head *BucketNode
}

// BucketNode is a node of singly linked list
type BucketNode struct {
	key  string
	next *BucketNode
}

// Insert will take key and inserts in hash table array if key is not present
func (hashTable *HashTable) Insert(key string) {
	index := Hash(key)
	hashTable.array[index].Insert(key)
}

// Search will take a key and returns true if found in hash table array
func (hashTable *HashTable) Search(key string) bool {
	index := Hash(key)
	return hashTable.array[index].Search(key)
}

// Delete will take a key and deletes from hash table array
func (hashTable *HashTable) Delete(key string) {
	index := Hash(key)
	hashTable.array[index].Delete(key)
}

// Init initializes the hash table array with the buckets
func (hashTable *HashTable) Init() {
	for i := 0; i < ArraySize; i++ {
		hashTable.array[i] = &Bucket{}
	}
}

// Insert will take a key, creates a node with the key and inserts it in the bucket
func (bucket *Bucket) Insert(key string) {
	if !bucket.Search(key) {
		newNode := &BucketNode{key: key}
		newNode.next = bucket.head
		bucket.head = newNode
	} else {
		fmt.Printf("Already present %s", key)
	}
}

// Search will take a key and returns true if bucket has the node with key
func (bucket *Bucket) Search(key string) bool {
	currentNode := bucket.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// Delete will take a key and deletes the ndoe with the key from bucket
func (bucket *Bucket) Delete(key string) {
	if bucket.head.key == key {
		bucket.head = bucket.head.next
		return
	}

	previousNode := bucket.head

	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
			return
		}
		previousNode = previousNode.next
	}
}

// Hash funtions takes key and returns int code greater than or equal to zero and less than ArraySize
func Hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}
