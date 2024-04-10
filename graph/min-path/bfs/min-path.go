package main

import "fmt"

type Edge struct {
	src    int
	dest   int
	weight int
}

type Graph struct {
	edges       []Edge
	adjacency   map[int][]Edge
	numVertices int
}

func NewGraph() *Graph {
	return &Graph{
		adjacency: make(map[int][]Edge),
	}
}

func (g *Graph) AddVertex(vertex int) {
	if _, exists := g.adjacency[vertex]; !exists {
		g.adjacency[vertex] = []Edge{}
		g.numVertices++
	}
}

func (g *Graph) AddEdge(src, dest, weight int) {
	edge := Edge{src, dest, weight}
	g.edges = append(g.edges, edge)
	g.adjacency[src] = append(g.adjacency[src], edge)                      // 将边添加到源顶点的邻接列表中
	g.adjacency[dest] = append(g.adjacency[dest], Edge{dest, src, weight}) // 将边添加到目标顶点的邻接列表中
}

func (g *Graph) GetEdges(vertex int) []Edge {
	return g.adjacency[vertex]
}
func (g *Graph) BFSMinPath(from, to int) int {
	queue := []int{from}
	distance := make([]int, g.numVertices)
	visited := make([]bool, g.numVertices)
	visited[from] = true

	for len(queue) != 0 {
		vertex := queue[0]
		queue = queue[1:]

		if vertex == to {
			return distance[to]
		}
		for _, edge := range g.GetEdges(vertex) {
			if !visited[edge.dest] {
				visited[edge.dest] = true
				distance[edge.dest] = distance[vertex] + 1
				queue = append(queue, edge.dest)
			}
		}
	}

	return -1
}
func main() {
	graph := NewGraph()

	// 添加顶点
	for i := 0; i < 6; i++ {
		graph.AddVertex(i)
	}
	// 添加边
	graph.AddEdge(0, 1, 1)
	graph.AddEdge(0, 2, 1)
	graph.AddEdge(0, 3, 1)
	graph.AddEdge(1, 4, 1)
	graph.AddEdge(1, 5, 1)
	graph.AddEdge(2, 4, 1)
	graph.AddEdge(3, 5, 1)
	minPath := graph.BFSMinPath(1, 5)
	fmt.Println(minPath)
}
