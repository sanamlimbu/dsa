package graph

import "errors"

var ErrKeyPresent = errors.New("key already present")
var ErrVertexMissing = errors.New("vertex not present")
var ErrEdgePresent = errors.New("edge already present")
var ErrInvalidWeight = errors.New("invalid weight")

type GraphRepresentation string

const (
	AdjacencyMatrixGraphRepresentation GraphRepresentation = "adjacency-matrix"
	AdjacencyListGraphRepresentation   GraphRepresentation = "adjacency-list"
)
