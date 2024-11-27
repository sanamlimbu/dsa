package binarysearchtree

import (
	"github.com/sanamlimbu/dsa"
)

type BinarySearchTree[T dsa.Ordered] struct {
	root *node[T]
}

// NewBinarySearchTree creates an empty binary tree.
func NewBinarySearchTree[T dsa.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

type node[T dsa.Ordered] struct {
	value T
	count int
	left  *node[T]
	right *node[T]
}

// Insert will add an element to the tree.
// If element if already present count is increased in the node.
func (t *BinarySearchTree[T]) Insert(element T) {
	if t.root == nil {
		t.root = &node[T]{value: element}
		return
	}

	t.root.insert(element)
}

func (n *node[T]) insert(element T) {
	switch {
	case n.value < element: // move right
		{
			if n.right == nil {
				n.right = &node[T]{value: element}
				break
			}

			n.right.insert(element)
		}
	case n.value > element: // move left
		{
			if n.left == nil {
				n.left = &node[T]{value: element}
				break
			}

			n.left.insert(element)
		}
	default: // same element
		n.count++
	}
}

// Search will take in an element and returns true if there is a node with that value else false.
// It also returns count of the element.
func (t *BinarySearchTree[T]) Search(element T) (bool, int) {
	if t.root == nil {
		return false, 0
	}

	return t.root.search(element)
}

func (n *node[T]) search(element T) (bool, int) {
	if n == nil {
		return false, 0
	}

	switch {
	case n.value < element: // move right
		return n.right.search(element)
	case n.value > element: // move left
		return n.left.search(element)
	default:
		return true, n.count
	}
}

// IsEmpty returns true if tree is empty. If not returns false.
func (t *BinarySearchTree[T]) IsEmpty() bool {
	return t.root == nil
}
