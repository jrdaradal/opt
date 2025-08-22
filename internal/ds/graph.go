package ds

import (
	"strings"

	"github.com/jrdaradal/opt/internal/fn"
)

type Vertex = string
type Edge [2]Vertex
type VertexSet = *Set[Vertex]
type EdgeSet = *Set[Edge]

const edgeGlue string = "-"

type Graph struct {
	Vertices []Vertex
	Edges    []Edge
	// Private fields
	index            map[Vertex]int
	connectedEdges   map[Vertex][]Edge
	adjacentVertices map[Vertex]VertexSet
}

func NewEdge(edge string) Edge {
	parts := strings.Split(edge, edgeGlue)
	return Edge{parts[0], parts[1]}
}

func (e Edge) String() string {
	return e[0] + edgeGlue + e[1]
}

func (e Edge) Tuple() (Vertex, Vertex) {
	return e[0], e[1]
}

func NewGraph(vertices string, edgePairs string) *Graph {
	g := &Graph{}
	g.Vertices = strings.Fields(vertices)
	g.index = make(map[Vertex]int)
	for i, vertex := range g.Vertices {
		g.index[vertex] = i
	}
	g.Edges = make([]Edge, 0)
	g.connectedEdges = make(map[Vertex][]Edge)
	g.adjacentVertices = make(map[Vertex]VertexSet)
	for _, edgePair := range strings.Fields(edgePairs) {
		edge := NewEdge(edgePair)
		v1, v2 := edge.Tuple()
		g.Edges = append(g.Edges, edge)
		g.addConnectedEdge(v1, edge)
		g.addConnectedEdge(v2, edge)
		g.addAdjacentVertex(v1, v2)
		g.addAdjacentVertex(v2, v1)
	}
	return g
}

func (g *Graph) addConnectedEdge(vertex Vertex, edge Edge) {
	g.connectedEdges[vertex] = append(g.connectedEdges[vertex], edge)
}

func (g *Graph) addAdjacentVertex(vertex Vertex, neighbor Vertex) {
	if _, ok := g.adjacentVertices[vertex]; !ok {
		g.adjacentVertices[vertex] = NewSet[Vertex]()
	}
	g.adjacentVertices[vertex].Add(neighbor)
}

func (g Graph) IndexOf(vertex Vertex) int {
	index, ok := g.index[vertex]
	return fn.Ternary(ok, index, -1)
}

func (g Graph) ConnectedEdges(vertex Vertex) []Edge {
	edges, ok := g.connectedEdges[vertex]
	return fn.Ternary(ok, edges, []Edge{})
}

func (g Graph) AdjacentVertices(vertex Vertex) []Vertex {
	neighbors, ok := g.adjacentVertices[vertex]
	if !ok {
		return []Vertex{}
	}
	return neighbors.Items()
}

func (g Graph) IsClique(vertices []Vertex) bool {
	vertexSet := SetFrom(vertices)
	for _, vertex := range vertices {
		adjacent := SetFrom(g.AdjacentVertices(vertex))
		adjacent.Add(vertex)
		if !vertexSet.Diff(adjacent).IsEmpty() {
			return false
		}
	}
	return true
}
