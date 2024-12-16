package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tr := NewTire()

	word1 := "hello"
	tr.Insert(word1)

	if !tr.Search(word1) {
		t.Errorf("expected word %s to be found in trie", word1)
	}

	word2 := "world"
	if tr.Search(word2) {
		t.Errorf("expected word %s to not be found in trie", word2)
	}

	tr.Insert(word2)

	if !tr.Search(word2) {
		t.Errorf("expected word %s to be found in trie after insertion", word2)
	}

	partialWord := "hell"
	if tr.Search(partialWord) {
		t.Errorf("expected partial word %s to not be found in trie", partialWord)
	}
}
