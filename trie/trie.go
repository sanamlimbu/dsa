package trie

// AlphabetSize is number of possible characters in trie.
// a-z characters are allowed.
const AlphabetSize = 26

// Node represents each node in trie.
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has pointer to the root node.
// This trie works for a-z alphabets only.
type Trie struct {
	root *Node
}

// NewTrie will create new trie.
func NewTire() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to trie.
func (t *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}

		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true
}

// Search will take in a word and returns true if that word is included in the trie.
func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'

		if currentNode.children[charIndex] == nil {
			return false
		}

		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}
