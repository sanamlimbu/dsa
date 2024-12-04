package graph

import (
	"fmt"

	s "github.com/sanamlimbu/dsa/stack"
)

// Depth-first search (DFS) is an algorithm for traversing or searching tree or graph data structures.
// The algorithm starts at the root node (selecting some arbitrary node as the root node in the case of a graph)
// and explores as far as possible along each branch before backtracking.
// Extra memory, usually a stack, is needed to keep track of the nodes discovered so far along a specified branch which helps in backtracking of the graph.
//
// The DFS algorithm works as follows:
//
// 1. Start by putting any one of the graph's vertices on top of a stack.
//
// 2. Take the top item of the stack and add it to the visited list.
//
// 3. Create a list of that vertex's adjacent nodes. Add the ones which aren't in the visited list to the top of the stack.
//
// 4. Keep repeating steps 2 and 3 until the stack is empty.
func DepthFirstSearch[T comparable](graph Graph[T], key T) {
	start := graph.GetVertex(key)
	if start == nil {
		fmt.Println("start vertex not found")
		return
	}

	vertices := graph.Vertices()

	stack := s.NewStack[*vertex[T]](s.SliceStackImplementation)
	stack.Push(start)

	visitedVertices := make(map[T]bool, len(vertices))
	visitedVertices[start.key] = true

	for !stack.IsEmpty() {
		vertex, err := stack.Pop()
		if err != nil {
			return
		}

		fmt.Println("current vertex:", vertex.key)

		for _, adj := range vertex.adjacent {
			if !visitedVertices[adj.key] {
				visitedVertices[adj.key] = true
				stack.Push(adj)
			}
		}
	}
}

func DepthFirstSearchByRecursion[T comparable](graph Graph[T], key T) {
	start := graph.GetVertex(key)
	if start == nil {
		fmt.Println("start vertex not found")
		return
	}

	vertices := graph.Vertices()

	visitedVertices := make(map[T]bool, len(vertices))

	dfs(start, visitedVertices)
}

func dfs[T comparable](vertex *vertex[T], visitedVertices map[T]bool) {
	visitedVertices[vertex.key] = true

	fmt.Println("current vertex:", vertex.key)

	for _, adj := range vertex.adjacent {
		if visitedVertices[adj.key] {
			dfs[T](adj, visitedVertices)
		}
	}
}
