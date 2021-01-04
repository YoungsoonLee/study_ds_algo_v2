package main

import (
	"container/heap"
	"fmt"
)

type Vertex int

type PriorityQueue struct {
	items []Vertex
	m     map[Vertex]int //value to index
	pr    map[Vertex]int // value to priority
}

func (pq *PriorityQueue) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.pr[pq.items[i]] < pq.pr[pq.items[j]]
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.m[pq.items[i]] = i
	pq.m[pq.items[j]] = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(pq.items)
	item := x.(Vertex)
	pq.m[item] = n
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	pq.m[item] = -1
	pq.items = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item Vertex, priority int) {
	pq.pr[item] = priority
	heap.Fix(pq, pq.m[item])
}

func (pq *PriorityQueue) addWithPriority(item Vertex, priority int) {
	heap.Push(pq, item)
	pq.update(item, priority)
}

const (
	Infinity      = int(^uint(0) >> 1) //maximum integer
	Uninitialized = -1                 //
)

func Dijkstra(g Graph, source Vertex) (distance map[Vertex]int, previous map[Vertex]Vertex) {
	distance = make(map[Vertex]int)
	previous = make(map[Vertex]Vertex)
	distance[source] = 0
	q := &PriorityQueue{[]Vertex{}, make(map[Vertex]int), make(map[Vertex]int)}
	for _, v := range g.Vertices() {
		if v != source {
			distance[v] = Infinity
		}
		previous[v] = Uninitialized
		q.addWithPriority(v, distance[v])
	}

	for len(q.items) != 0 {
		u := heap.Pop(q).(Vertex)
		for _, v := range g.Neighbors(u) {
			alt := distance[u] + g.Weight(u, v)
			if alt < distance[v] {
				distance[v] = alt
				previous[v] = u
				q.update(v, alt)
			}
		}
	}
	return distance, previous
}

type Graph interface {
	Vertices() []Vertex
	Neighbors(v Vertex) []Vertex
	Weight(u, v Vertex) int
}

type AdjacencyMap struct {
	ids   map[string]Vertex
	names map[Vertex]string
	edges map[Vertex]map[Vertex]int
}

func NewGraph(ids map[string]Vertex) AdjacencyMap {
	g := AdjacencyMap{ids: ids}
	g.names = make(map[Vertex]string)
	for k, v := range ids {
		g.names[v] = k
	}
	g.edges = make(map[Vertex]map[Vertex]int)
	return g
}

func (g AdjacencyMap) AddEdge(u, v string, w int) {
	if _, ok := g.edges[g.ids[u]]; !ok {
		g.edges[g.ids[u]] = make(map[Vertex]int)
	}
	g.edges[g.ids[u]][g.ids[v]] = w
}

func (g AdjacencyMap) GetPath(v Vertex, previous map[Vertex]Vertex) (path string) {
	path = g.names[v]
	for previous[v] >= 0 {
		v = previous[v]
		path = g.names[v] + path
	}
	return path
}

func (g AdjacencyMap) Vertices() (vertices []Vertex) {
	for _, v := range g.ids {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g AdjacencyMap) Neighbors(u Vertex) (vertices []Vertex) {
	for v := range g.edges[u] {
		vertices = append(vertices, v)
	}
	return vertices
}

func (g AdjacencyMap) Weight(u, v Vertex) int {
	return g.edges[u][v]
}

func main() {
	G := NewGraph(map[string]Vertex{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6,
	})

	G.AddEdge("a", "b", 7)
	G.AddEdge("a", "c", 9)
	G.AddEdge("a", "f", 14)
	G.AddEdge("b", "c", 10)
	G.AddEdge("b", "d", 15)
	G.AddEdge("c", "d", 11)
	G.AddEdge("c", "f", 2)
	G.AddEdge("d", "e", 6)
	G.AddEdge("e", "f", 9)

	distance, previous := Dijkstra(G, G.ids["a"])

	fmt.Printf("Distance to %s is %d, Path: %s\n", "e", distance[G.ids["e"]], G.GetPath(G.ids["e"], previous))
	fmt.Printf("Distance to %s is %d, Path: %s\n", "f", distance[G.ids["f"]], G.GetPath(G.ids["f"], previous))
}
