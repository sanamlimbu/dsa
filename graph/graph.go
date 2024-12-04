package graph

import "errors"

var ErrKeyPresent = errors.New("key already present")
var ErrVertexMissing = errors.New("vertex not present")
var ErrEdgePresent = errors.New("edge already present")

type vertex[T comparable] struct {
	key      T
	adjacent []*vertex[T]
}

type Graph[T comparable] interface {
	AddVertex(key T) error
	AddEdge(from, to T) error
	Vertices() []*vertex[T]
	GetVertex(key T) *vertex[T]
}
