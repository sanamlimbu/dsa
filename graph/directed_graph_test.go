package graph

import (
	"testing"
)

func TestDirectedGraph(t *testing.T) {
	g := NewDirectedGraph[string]()

	t.Run("Add vertex", func(t *testing.T) {
		err := g.AddVertex("A")
		if err != nil {
			t.Errorf("unexpected error adding vertex: %v", err)
		}

		err = g.AddVertex("A")
		if err != ErrKeyPresent {
			t.Errorf("expected ErrKeyPresent, got: %v", err)
		}
	})

	t.Run("Add edge", func(t *testing.T) {
		err := g.AddVertex("B")
		if err != nil {
			t.Errorf("unexpected error adding vertex: %v", err)
		}

		err = g.AddEdge("A", "B")
		if err != nil {
			t.Errorf("unexpected error adding edge: %v", err)
		}

		err = g.AddEdge("A", "B")
		if err != ErrEdgePresent {
			t.Errorf("expected ErrEdgePresent, got: %v", err)
		}

		err = g.AddEdge("A", "C")
		if err != ErrVertexMissing {
			t.Errorf("expected ErrVertexMissing, got: %v", err)
		}
	})
}
