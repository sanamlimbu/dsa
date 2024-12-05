package graph

import (
	"io"
	"os"
	"slices"
	"testing"
)

// captureOutput captures output from the console and returns it.
func captureOutput(f func()) string {
	storeStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()
	w.Close()

	out, _ := io.ReadAll(r)

	// Restore the stdout
	os.Stdout = storeStdout

	return string(out)
}

var expectedDfsOutputs []string = []string{
	"current vertex: 0\ncurrent vertex: 1\ncurrent vertex: 2\ncurrent vertex: 4\ncurrent vertex: 3\n",
	"current vertex: 0\ncurrent vertex: 2\ncurrent vertex: 4\ncurrent vertex: 1\ncurrent vertex: 3\n",
	"current vertex: 0\ncurrent vertex: 2\ncurrent vertex: 1\ncurrent vertex: 4\ncurrent vertex: 3\n",
	"current vertex: 0\ncurrent vertex: 3\ncurrent vertex: 1\ncurrent vertex: 2\ncurrent vertex: 4\n",
	"current vertex: 0\ncurrent vertex: 3\ncurrent vertex: 2\ncurrent vertex: 4\ncurrent vertex: 1\n",
	"current vertex: 0\ncurrent vertex: 3\ncurrent vertex: 2\ncurrent vertex: 1\ncurrent vertex: 4\n",
}

func TestDepthFirstSearch(t *testing.T) {
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
		DepthFirstSearch[int](graph, 0)
	})

	if !slices.Contains(expectedDfsOutputs, output) {
		t.Errorf("DFS output was incorrect, got: %q", output)
	}
}

func TestDepthFirstSearchByRecursion(t *testing.T) {
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
		DepthFirstSearchByRecursion[int](graph, 0)
	})

	if !slices.Contains(expectedDfsOutputs, output) {
		t.Errorf("DFS output was incorrect, got: %q", output)
	}
}
