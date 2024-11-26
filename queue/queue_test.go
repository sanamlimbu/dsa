package queue

import "testing"

func TestLinkedListQueue(t *testing.T) {
	queue := NewQueue[int](LinkedListQueueImplementation)

	if !queue.IsEmpty() {
		t.Error("expected queue to be empty")
	}

	queue.Enqueue(1)
	queue.Enqueue(2)

	if queue.IsEmpty() {
		t.Error("expected queue not to be empty")
	}

	if val, _ := queue.Peek(); val != 1 {
		t.Errorf("expected front element to be 1, got %d", val)
	}

	if val, _ := queue.Dequeue(); val != 1 {
		t.Errorf("expected dequeued element to be 1, got %d", val)
	}

	if queue.IsEmpty() {
		t.Error("expected queue not to be empty")
	}

	if val, _ := queue.Peek(); val != 2 {
		t.Errorf("expected front element to be 2, got %d", val)
	}

	if val, _ := queue.Dequeue(); val != 2 {
		t.Errorf("expected dequeued element to be 2, got %d", val)
	}

	if !queue.IsEmpty() {
		t.Error("expected queue to be empty")
	}
}

func TestSliceQueue(t *testing.T) {
	queue := NewQueue[int](SliceQueueImplementation)

	if !queue.IsEmpty() {
		t.Error("expected queue to be empty")
	}

	queue.Enqueue(1)
	queue.Enqueue(2)

	if queue.IsEmpty() {
		t.Error("expected queue not to be empty")
	}

	if val, _ := queue.Peek(); val != 1 {
		t.Errorf("expected front element to be 1, got %d", val)
	}

	if val, _ := queue.Dequeue(); val != 1 {
		t.Errorf("expected dequeued element to be 1, got %d", val)
	}

	if queue.IsEmpty() {
		t.Error("expected queue not to be empty")
	}

	if val, _ := queue.Peek(); val != 2 {
		t.Errorf("expected front element to be 2, got %d", val)
	}

	if val, _ := queue.Dequeue(); val != 2 {
		t.Errorf("expected dequeued element to be 2, got %d", val)
	}

	if !queue.IsEmpty() {
		t.Error("expected queue to be empty")
	}
}
