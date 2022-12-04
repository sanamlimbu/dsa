package binarysearchtree

type Node struct {
	key   int
	left  *Node
	right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(key int) {
	if n.key < key {
		// move right
		if n.right == nil {
			n.right = &Node{key: key}
		} else {
			n.right.Insert(key)
		}
	} else if n.key > key {
		// move left
		if n.left == nil {
			n.left = &Node{key: key}
		} else {
			n.left.Insert(key)
		}
	}
}

// Search will take in a key value
// and return true if there is a node with that value else false
func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if n.key < key {
		// move right
		return n.right.Search(key)
	} else if n.key > key {
		// move left
		return n.left.Search(key)
	}

	return true
}
