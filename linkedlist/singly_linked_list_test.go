package linkedlist

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int, int]()

	list.Insert(12, 12, 0)
	list.Insert(10, 10, 1)
	list.Insert(20, 20, 2)
	list.Insert(30, 30, 2)

	current := list.Slice()
	expected := []pair[int, int]{
		{
			key: 10, value: 10,
		},
		{
			key: 30, value: 30,
		},
		{
			key: 20, value: 20,
		}, {
			key: 12, value: 12,
		},
	}

	if !reflect.DeepEqual(current, expected) {
		t.Errorf("expected: %v, got: %v", expected, current)
	}

	if got := list.Search(10); !got {
		t.Errorf("Search(10) = %v, want true", got)
	}

	if got := list.Search(40); got {
		t.Errorf("Search(40) = %v, want false", got)
	}

	list.Delete(2)

	if got := list.Search(30); got {
		t.Errorf("Search(30) after Delete(2) = %v, want false", got)
	}

	current = list.Slice()
	expected = []pair[int, int]{
		{
			key: 10, value: 10,
		},
		{
			key: 20, value: 20,
		}, {
			key: 12, value: 12,
		},
	}

	if !reflect.DeepEqual(current, expected) {
		t.Errorf("expected: %v, got: %v", expected, current)
	}

	length := list.Length()

	if length != 3 {
		t.Errorf("expected: 3, got: %d", length)
	}

	list.DeleteKey(12)

	current = list.Slice()
	expected = []pair[int, int]{
		{
			key: 10, value: 10,
		},
		{
			key: 20, value: 20,
		},
	}

	if !reflect.DeepEqual(current, expected) {
		t.Errorf("expected: %v, got: %v", expected, current)
	}
}
