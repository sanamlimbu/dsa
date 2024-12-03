package graph

import (
	"errors"
)

var ErrKeyPresent = errors.New("key already present")
var ErrVertexMissing = errors.New("vertex not present")
var ErrEdgePresent = errors.New("edge already present")

// DirectedGraph represents an adjacency list graph.
type DirectedGraph[T comparable] struct {
	vertices []*vertex[T]
}

func NewDirectedGraph[T comparable]() *DirectedGraph[T] {
	return &DirectedGraph[T]{}
}

type vertex[T comparable] struct {
	key      T
	adjacent []*vertex[T]
}

// AddVertex adds a vertex to the graph.
// Returns error if key is already present.
func (g *DirectedGraph[T]) AddVertex(key T) error {
	if contains(g.vertices, key) {
		return ErrKeyPresent
	}

	g.vertices = append(g.vertices, &vertex[T]{key: key})

	return nil
}

// AddEdge adds an edge to the graph.
func (g *DirectedGraph[T]) AddEdge(from, to T) error {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		return ErrVertexMissing
	}

	// Check for duplicate edge
	if contains(fromVertex.adjacent, to) {
		return ErrEdgePresent
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)

	return nil
}

// getVertex returns a pointer to the vertex with given key.
// If key not found, nil is returned.
func (g *DirectedGraph[T]) getVertex(key T) *vertex[T] {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}

	return nil
}

// contains checks if the vertex list has given key value.
func contains[T comparable](vertices []*vertex[T], key T) bool {
	for _, v := range vertices {
		if v.key == key {
			return true
		}
	}

	return false
}
