package graph

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	key       int
	adjacenct []*Vertex
}

// AddVertex adds a Vertex to the Graph
func (g *Graph) AddVertex(key int) error {
	if contains(g.vertices, key) {
		return fmt.Errorf("key already present")
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
	return nil
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) error {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	// no vetex ?
	if fromVertex == nil || toVertex == nil {
		return fmt.Errorf("vertex not present")
	}

	// duplicate edge ?
	if contains(fromVertex.adjacenct, to) {
		return fmt.Errorf("edge already present")
	}

	// add edge
	fromVertex.adjacenct = append(fromVertex.adjacenct, toVertex)

	return nil
}

// getVertex returns a pointer to the Vertex whith a key integer
func (g *Graph) getVertex(key int) *Vertex {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}
	return nil
}

// contains checks if the Vertex list has given key value
func contains(vertices []*Vertex, key int) bool {
	for _, v := range vertices {
		if v.key == key {
			return true
		}
	}
	return false
}

// AddEdge adds a

// Print will print the adjacent list of each vertex of graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacenct {
			fmt.Printf(" %v ", v.key)
		}
	}
}
