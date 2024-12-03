package linkedlist

import (
	"reflect"
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	list := NewSinglyLinkedList[int, int]()

	list.InsertAtPosition(12, 12, 0)
	list.InsertAtPosition(10, 10, 1)
	list.InsertAtPosition(20, 20, 2)
	list.InsertAtPosition(30, 30, 2)

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

	if _, got := list.Search(10); !got {
		t.Errorf("Search(10) = %v, want true", got)
	}

	if _, got := list.Search(40); got {
		t.Errorf("Search(40) = %v, want false", got)
	}

	list.Delete(2)

	if _, got := list.Search(30); got {
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

	list.Insert(50, 50)
	list.InsertAtPosition(5, 5, 1)

	current = list.Slice()
	expected = []pair[int, int]{
		{
			key: 5, value: 5,
		},
		{
			key: 10, value: 10,
		},
		{
			key: 20, value: 20,
		},
		{
			key: 50, value: 50,
		},
	}

	if !reflect.DeepEqual(current, expected) {
		t.Errorf("expected: %v, got: %v", expected, current)
	}
}
