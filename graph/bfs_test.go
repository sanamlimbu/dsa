package graph

import (
	"slices"
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	var expectedBfsOutputs []string = []string{
		"current vertex: 0\ncurrent vertex: 1\ncurrent vertex: 2\ncurrent vertex: 3\ncurrent vertex: 4\n",
		"current vertex: 0\ncurrent vertex: 1\ncurrent vertex: 3\ncurrent vertex: 2\ncurrent vertex: 4\n",
		"current vertex: 0\ncurrent vertex: 2\ncurrent vertex: 1\ncurrent vertex: 3\ncurrent vertex: 4\n",
		"current vertex: 0\ncurrent vertex: 2\ncurrent vertex: 3\ncurrent vertex: 1\ncurrent vertex: 4\n",
		"current vertex: 0\ncurrent vertex: 3\ncurrent vertex: 1\ncurrent vertex: 2\ncurrent vertex: 4\n",
		"current vertex: 0\ncurrent vertex: 3\ncurrent vertex: 2\ncurrent vertex: 1\ncurrent vertex: 4\n",
	}

	graph := NewDirectedUnweightedGraphAsList[int]()
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 3)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 4)

	output := captureOutput(func() {
		BreadthFirstSearch(graph, 0)
	})

	if !slices.Contains(expectedBfsOutputs, output) {
		t.Errorf("DFS output was incorrect, got: %q", output)
	}
}
