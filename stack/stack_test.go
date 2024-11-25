package stack

import (
	"testing"
)

func TestSliceStack(t *testing.T) {
	stack := NewSliceStack[int]()

	// Test IsEmpty on an empty stack
	if !stack.IsEmpty() {
		t.Errorf("expected stack to be empty")
	}

	// Test Push
	stack.Push(10)
	stack.Push(20)

	if stack.IsEmpty() {
		t.Errorf("expected stack to be non-empty after push operations")
	}

	// Test Peek
	top, err := stack.Peek()
	if err != nil {
		t.Errorf("unexpected error in Peek: %v", err)
	}
	if top != 20 {
		t.Errorf("expected top element to be 20, got %d", top)
	}

	// Test Pop
	element, err := stack.Pop()
	if err != nil {
		t.Errorf("unexpected error in Pop: %v", err)
	}
	if element != 20 {
		t.Errorf("expected popped element to be 20, got %d", element)
	}

	element, err = stack.Pop()
	if err != nil {
		t.Errorf("unexpected error in Pop: %v", err)
	}
	if element != 10 {
		t.Errorf("expected popped element to be 10, got %d", element)
	}

	// Test Pop on empty stack
	_, err = stack.Pop()
	if err == nil {
		t.Errorf("expected error when popping from empty stack")
	}
}

func TestLinkedListStack(t *testing.T) {
	stack := NewLinkedListStack[int]()

	// Test IsEmpty on an empty stack
	if !stack.IsEmpty() {
		t.Errorf("expected stack to be empty")
	}

	// Test Push
	stack.Push(10)
	stack.Push(20)

	if stack.IsEmpty() {
		t.Errorf("expected stack to be non-empty after push operations")
	}

	// Test Peek
	top, err := stack.Peek()
	if err != nil {
		t.Errorf("unexpected error in Peek: %v", err)
	}
	if top != 20 {
		t.Errorf("expected top element to be 20, got %d", top)
	}

	// Test Pop
	element, err := stack.Pop()
	if err != nil {
		t.Errorf("unexpected error in Pop: %v", err)
	}
	if element != 20 {
		t.Errorf("expected popped element to be 20, got %d", element)
	}

	element, err = stack.Pop()
	if err != nil {
		t.Errorf("unexpected error in Pop: %v", err)
	}
	if element != 10 {
		t.Errorf("expected popped element to be 10, got %d", element)
	}

	// Test Pop on empty stack
	_, err = stack.Pop()
	if err == nil {
		t.Errorf("expected error when popping from empty stack")
	}
}
