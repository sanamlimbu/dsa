package graph

import (
	"fmt"
	"math"
)

// Dijkstra's algorithm is an algorithm for finding the shortest paths between vertices in a weighted graph.
// Dijkstra's algorithm finds the shortest path from a given source vertex to every other vertex.
func Dijkstra[T comparable](graph *DirectedWeightedGraphAsMatrix[T], source T) {
	vertices := graph.Vertices()
	length := len(vertices)

	shortestDistances := make(map[T]int, length)
	completedVertices := make(map[T]bool, length)

	infinity := math.MaxInt - 1

	// Make all vertices distances infinity.
	for _, v := range vertices {
		shortestDistances[v] = infinity
	}

	shortestDistances[source] = 0

	for i := 0; i < length-1; i++ {
		vertex := minimumDistanceVertex(shortestDistances, completedVertices)
		completedVertices[vertex] = true

		// Update distances of adjacent vertices of picked vertex.
		for j := 0; j < length; j++ {
			adjacent := vertices[j]
			weight := graph.Weight(vertex, adjacent)

			if !completedVertices[adjacent] &&
				graph.IsEdge(vertex, adjacent) &&
				shortestDistances[vertex] != infinity &&
				shortestDistances[vertex]+weight < shortestDistances[adjacent] {
				shortestDistances[adjacent] = shortestDistances[vertex] + weight
			}
		}
	}

	// Print solution.
	for key, value := range shortestDistances {
		fmt.Println("Vertex", key, ":", value)
	}
}

// minimumDistanceVertex returns minimum distance vertex from the set of vertices not yet processed.
func minimumDistanceVertex[T comparable](shortestDistances map[T]int, completedVertices map[T]bool) T {
	min := math.MaxInt
	var vertex T

	for key, value := range shortestDistances {
		if value <= min && !completedVertices[key] {
			min = value
			vertex = key
		}
	}

	return vertex
}
