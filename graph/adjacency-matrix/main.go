package main

import (
	"errors"
	"fmt"
)

type AdjacencyMatrix struct {
	Vertices  int //nummber off vertices
	Edges     int
	GraphType GraphType
	AdjMatrix [][]int
}

type GraphType string

const (
	DIRECTED   GraphType = "DIRECTED"
	UNDIRECTED GraphType = "UNDIRECTED"
)

type Graph interface {
	Init()
	AddEdge(vertexOne int, vertexTwo int) error
	AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int) error
	RemoveEdge(vertexOne int, vertexTwo int) error
	HasEdge(vertexOne int, vertexTwo int) bool
	GetGraphType() GraphType
	GetAdjacentNodesForVertex(vertex int) map[int]bool
	GetWeightOfEdge(vertexOne int, vertexTwo int) (int, error)
	GetNumberOfVertices() int
	GetNumberOfEdges() int
	GetIndegreeForVertex(vertex int) int
}

func (g *AdjacencyMatrix) Init() {
	g.AdjMatrix = make([][]int, g.Vertices)
	g.Edges = 0
	for i := 0; i < g.Vertices; i++ {
		g.AdjMatrix[i] = make([]int, g.Vertices) // default initialization is 0
	}
}

func (g *AdjacencyMatrix) AddEdge(vertexOne int, vertexTwo int) error {
	if vertexOne >= g.Vertices || vertexTwo >= g.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}

	g.AdjMatrix[vertexOne][vertexTwo] = 1
	g.Edges++
	if g.GraphType == UNDIRECTED {
		g.AdjMatrix[vertexTwo][vertexOne] = 1
		g.Edges++
	}
	return nil
}

func (g *AdjacencyMatrix) AddEdgeWithWeight(vertexOne, vertexTwo, weight int) error {
	if vertexOne >= g.Vertices || vertexTwo >= g.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	g.AdjMatrix[vertexOne][vertexTwo] = weight
	if g.GraphType == UNDIRECTED {
		g.AdjMatrix[vertexTwo][vertexOne] = weight
		g.Edges++
	}
	return nil
}

func (g *AdjacencyMatrix) RemoveEdge(vertexOne, vertexTwo int) error {
	if vertexOne >= g.Vertices || vertexTwo >= g.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return errors.New("Index out of bounds")
	}
	g.AdjMatrix[vertexOne][vertexTwo] = 0
	g.Edges--
	if g.GraphType == UNDIRECTED {
		g.AdjMatrix[vertexTwo][vertexOne] = 0
		g.Edges--
	}
	return nil
}

func (g *AdjacencyMatrix) HasEdge(vertexOne, vertexTwo int) bool {
	if vertexOne >= g.Vertices || vertexTwo >= g.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return false
	}
	return g.AdjMatrix[vertexOne][vertexTwo] != 0
}

func (g *AdjacencyMatrix) GetGraphType() GraphType {
	return g.GraphType
}

func (g AdjacencyMatrix) GetAdjacentNodesForVertex(vertex int) map[int]bool {
	adjacencyMatrixVertices := map[int]bool{}
	if vertex >= g.Vertices || vertex < 0 {
		return adjacencyMatrixVertices
	}
	for i := 0; i < g.Vertices; i++ {
		if g.AdjMatrix[vertex][i] != 0 {
			adjacencyMatrixVertices[i] = (g.AdjMatrix[vertex][i] != 0)
		}
	}
	return adjacencyMatrixVertices
}

func (g *AdjacencyMatrix) GetWeightOfEdge(vertexOne, vertexTwo int) (int, error) {
	if vertexOne >= g.Vertices || vertexTwo >= g.Vertices || vertexOne < 0 || vertexTwo < 0 {
		return 0, errors.New("Error getting weight for vertex")
	}
	return g.AdjMatrix[vertexOne][vertexTwo], nil
}

func (g *AdjacencyMatrix) GetNumberOfVertices() int {
	return g.Vertices
}

func (g *AdjacencyMatrix) GetNumberOfEdges() int {
	return g.Edges
}

func (g *AdjacencyMatrix) GetIndegreeForVertex(vertex int) int {
	indegree := 0
	AdjacencyNodes := g.GetAdjacentNodesForVertex(vertex)
	for key := range AdjacencyNodes {
		if AdjacencyNodes[key] {
			indegree++
		}
	}
	return indegree
}

func main() {
	var testAdjMatrixDirected = &AdjacencyMatrix{4, 0, DIRECTED, nil}
	testAdjMatrixDirected.Init()
	err := testAdjMatrixDirected.AddEdge(2, 1)
	if err != nil {
		fmt.Printf("Error adding edge")
	}

	if testAdjMatrixDirected.AdjMatrix[2][1] != 1 {
		fmt.Printf("Data not found at index")
	}

	if testAdjMatrixDirected.AdjMatrix[1][2] != 0 {
		fmt.Printf("Data not found at index")
	}
}
