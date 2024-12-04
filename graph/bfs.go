package graph

import (
	"fmt"

	q "github.com/sanamlimbu/dsa/queue"
)

// Breadth First Search (BFS) is a fundamental graph traversal algorithm.
// It begins with a node, then first traverses all its adjacent. Once all adjacent are visited, then their adjacent are traversed.
// The only catch here is that, unlike trees, graphs may contain cycles, so we may come to the same node again.
// To avoid processing a node more than once, we use a boolean visited array.
//
//  1. Initialization: Enqueue the given source vertex into a queue and mark it as visited.
//  2. Exploration: While the queue is not empty:
//     2.1 Dequeue a node from the queue and visit it (e.g., print its value).
//     2.2 For each unvisited neighbor of the dequeued node: Enqueue the neighbor into the queue.
//     2.3 Mark the neighbor as visited.
//  3. Termination: Repeat step 2 until the queue is empty.
func BreadthFirstSearch[T comparable](graph Graph[T], key T) {
	start := graph.GetVertex(key)
	if start == nil {
		fmt.Println("start vertex not found")
		return
	}

	vertices := graph.Vertices()

	queue := q.NewQueue[*vertex[T]](q.SliceQueueImplementation)
	queue.Enqueue(start)

	visitedVertices := make(map[T]bool, len(vertices))
	visitedVertices[key] = true

	for !queue.IsEmpty() {
		vertex, err := queue.Dequeue()
		if err != nil {
			return
		}

		fmt.Println("current vertex:", vertex.key)

		for _, adj := range vertex.adjacent {
			if !visitedVertices[adj.key] {
				visitedVertices[adj.key] = true
				queue.Enqueue(adj)
			}
		}
	}
}
