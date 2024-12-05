package graph

type weightedGraphVertex[T comparable] struct {
	key      T
	adjacent []*weightedGraphEdge[T]
}

type weightedGraphEdge[T comparable] struct {
	weight int
	vertex *weightedGraphVertex[T]
}

type DirectedWeightedGraph[T comparable] interface {
	AddVertex(key T) error
	AddEdge(to, from T, weight int) error
	IsEmpty() bool
}

func NewDirectedWeightedGraph[T comparable](representation GraphRepresentation) DirectedWeightedGraph[T] {
	if representation == AdjacencyMatrixGraphRepresentation {
		return NewDirectedWeightedGraphAsMatrix[T]()
	}

	return NewDirectedWeightedGraphAsList[T]()
}

// Directed weighted graph represented as adjacency matrix.
type DirectedWeightedGraphAsMatrix[T comparable] struct {
	matrix         [][]int
	vertexIndexMap map[T]int
}

func NewDirectedWeightedGraphAsMatrix[T comparable]() *DirectedWeightedGraphAsMatrix[T] {
	return &DirectedWeightedGraphAsMatrix[T]{
		matrix:         make([][]int, 0),
		vertexIndexMap: make(map[T]int, 0),
	}
}

func (g *DirectedWeightedGraphAsMatrix[T]) IsEmpty() bool {
	return len(g.matrix) == 0
}

// AddVertex adds a vertex to the graph.
// Returns error if key is already present.
func (g *DirectedWeightedGraphAsMatrix[T]) AddVertex(key T) error {
	_, found := g.vertexIndexMap[key]
	if found {
		return ErrKeyPresent
	}

	length := len(g.matrix)

	g.vertexIndexMap[key] = length

	for i := range g.matrix {
		g.matrix[i] = append(g.matrix[i], 0)
	}

	row := make([]int, length+1)

	g.matrix = append(g.matrix, row)

	return nil
}

// AddEdge adds an edge to the graph with given weight.
func (g *DirectedWeightedGraphAsMatrix[T]) AddEdge(from, to T, weight int) error {
	if weight == 0 {
		return ErrInvalidWeight
	}

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
	// Weight 0 means there is no edge.
	if g.matrix[fromIndex][toIndex] != 0 {
		return ErrEdgePresent
	}

	g.matrix[fromIndex][toIndex] = weight

	return nil
}

// Directed weighted graph represented as adjacency list.
type DirectedWeightedGraphAsList[T comparable] struct {
	vertices []*weightedGraphVertex[T]
}

func NewDirectedWeightedGraphAsList[T comparable]() *DirectedWeightedGraphAsList[T] {
	return &DirectedWeightedGraphAsList[T]{}
}

func (g *DirectedWeightedGraphAsList[T]) IsEmpty() bool {
	return len(g.vertices) == 0
}

// AddVertex adds a vertex to the graph.
// Returns error if key is already present.
func (g *DirectedWeightedGraphAsList[T]) AddVertex(key T) error {
	// Check if there is already a vertex.
	for _, v := range g.vertices {
		if v.key == key {
			return ErrKeyPresent
		}
	}

	g.vertices = append(g.vertices, &weightedGraphVertex[T]{key: key})

	return nil
}

// AddEdge adds an edge to the graph with given weight.
func (g *DirectedWeightedGraphAsList[T]) AddEdge(from, to T, weight int) error {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex == nil || toVertex == nil {
		return ErrVertexMissing
	}

	// Check for duplicate edge.
	for _, edge := range fromVertex.adjacent {
		if edge.vertex.key == to {
			return ErrEdgePresent
		}
	}

	fromVertex.adjacent = append(fromVertex.adjacent, &weightedGraphEdge[T]{
		weight: weight,
		vertex: toVertex,
	})

	return nil
}

// GetVertex returns a pointer to the vertex with given key.
// If key not found, nil is returned.
func (g *DirectedWeightedGraphAsList[T]) GetVertex(key T) *weightedGraphVertex[T] {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}

	return nil
}

// Vertices returns copy of vertices of the graph.
func (g *DirectedWeightedGraphAsList[T]) Vertices() []*weightedGraphVertex[T] {
	result := make([]*weightedGraphVertex[T], len(g.vertices))

	copy(result, g.vertices)

	return result
}
