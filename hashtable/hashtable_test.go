package hashtable

import (
	"testing"
)

func TestHashTable(t *testing.T) {
	ht := NewHashTable[string]()

	if _, found := ht.Get("key"); found {
		t.Errorf("get failed: did not expect any key to exist in empty hash table")
	}

	ht.Delete("key")

	ht.Set("key1", "value1")
	ht.Set("key2", "value2")
	ht.Set("key3", "value3")

	if _, found := ht.Get("key1"); !found {
		t.Errorf("get failed: expected key1 to exist")
	}

	if _, found := ht.Get("key2"); !found {
		t.Errorf("get failed: expected key2 to exist")
	}

	if _, found := ht.Get("key3"); !found {
		t.Errorf("get failed: expected key3 to exist")
	}

	if _, found := ht.Get("nonexistent"); found {
		t.Errorf("get failed: did not expect nonexistent key to exist")
	}

	ht.Set("key1", "new_value1")
	if value, _ := ht.Get("key1"); value != "new_value1" {
		t.Errorf("get failed: expected key1 to exist after update")
	}

	ht.Delete("key1")
	if _, found := ht.Get("key1"); found {
		t.Errorf("delete failed: expected key1 to be deleted")
	}

	ht.Delete("nonexistent")
}
