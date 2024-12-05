package graph

type vertex[T comparable] struct {
	key      T
	adjacent []*vertex[T]
}

type DirectedUnweightedGraph[T comparable] interface {
	AddVertex(key T) error
	AddEdge(to, from T) error
	IsEmpty() bool
}

func NewDirectedUnweightedGraph[T comparable](representation GraphRepresentation) DirectedUnweightedGraph[T] {
	if representation == AdjacencyMatrixGraphRepresentation {
		return NewDirectedUnweightedGraphAsMatrix[T]()
	}

	return NewDirectedUnweightedGraphAsList[T]()
}

// Directed unweighted graph represented as adjacency list.
type DirectedUnweightedGraphAsList[T comparable] struct {
	vertices []*vertex[T]
}

func NewDirectedUnweightedGraphAsList[T comparable]() *DirectedUnweightedGraphAsList[T] {
	return &DirectedUnweightedGraphAsList[T]{}
}

func (g *DirectedUnweightedGraphAsList[T]) IsEmpty() bool {
	return len(g.vertices) == 0
}

// AddVertex adds a vertex to the graph.
// Returns error if key is already present.
func (g *DirectedUnweightedGraphAsList[T]) AddVertex(key T) error {
	if contains(g.vertices, key) {
		return ErrKeyPresent
	}

	g.vertices = append(g.vertices, &vertex[T]{key: key})

	return nil
}

// AddEdge adds an edge to the graph.
func (g *DirectedUnweightedGraphAsList[T]) AddEdge(from, to T) error {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

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

// Vertices returns copy of vertices of the graph.
func (g *DirectedUnweightedGraphAsList[T]) Vertices() []*vertex[T] {
	result := make([]*vertex[T], len(g.vertices))

	copy(result, g.vertices)

	return result
}

// GetVertex returns a pointer to the vertex with given key.
// If key not found, nil is returned.
func (g *DirectedUnweightedGraphAsList[T]) GetVertex(key T) *vertex[T] {
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

// Directed unweighted graph represented as adjacency matrix.
type DirectedUnweightedGraphAsMatrix[T comparable] struct {
	matrix         [][]bool
	vertexIndexMap map[T]int
}

func NewDirectedUnweightedGraphAsMatrix[T comparable]() *DirectedUnweightedGraphAsMatrix[T] {
	return &DirectedUnweightedGraphAsMatrix[T]{
		matrix:         make([][]bool, 0),
		vertexIndexMap: make(map[T]int, 0),
	}
}

func (g *DirectedUnweightedGraphAsMatrix[T]) IsEmpty() bool {
	return len(g.matrix) == 0
}

// AddVertex adds a vertex to the graph.
// Returns error if key is already present.
func (g *DirectedUnweightedGraphAsMatrix[T]) AddVertex(key T) error {
	_, found := g.vertexIndexMap[key]
	if found {
		return ErrKeyPresent
	}

	length := len(g.matrix)

	g.vertexIndexMap[key] = length

	for i := range g.matrix {
		g.matrix[i] = append(g.matrix[i], false)
	}

	row := make([]bool, length+1)

	g.matrix = append(g.matrix, row)

	return nil
}

// AddEdge adds an edge to the graph.
func (g *DirectedUnweightedGraphAsMatrix[T]) AddEdge(from, to T) error {
	// Row index
	fromIndex, found := g.vertexIndexMap[from]
	if !found {
		return ErrVertexMissing
	}

	// Column index
	toIndex, found := g.vertexIndexMap[to]
	if !found {
		return ErrVertexMissing
	}

	// Check for duplicate edge.
	// There is edge if cell is true.
	if g.matrix[fromIndex][toIndex] {
		return ErrEdgePresent
	}

	g.matrix[fromIndex][toIndex] = true

	return nil
}
