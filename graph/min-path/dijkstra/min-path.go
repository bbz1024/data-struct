package main

import (
	"fmt"
	"math"
)

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
func (g *Graph) DijkStraMinPath(from, to int) int {
	//初始化
	//标记各顶点是否已经找到最短路径
	final := make([]bool, g.numVertices)
	//最短路径长度
	dist := make([]int, g.numVertices)
	//路径的前驱
	fromPath := make([]int, g.numVertices)
	for i := 0; i < g.numVertices; i++ {
		dist[i] = math.MaxInt32
		fromPath[i] = -1
	}
	final[from] = true
	dist[from] = 0
	fromPath[from] = -1
	//初始化顶点
	//边
	for _, edge := range g.GetEdges(from) {
		dist[edge.dest] = edge.weight
		fromPath[edge.dest] = from
	}
	//第一轮：循环所有的结点，找到还没确定最短路径，且dist最小的顶点Vi,令final[i]=true
	for i := 0; i < g.numVertices-1; i++ {
		idx := -1
		min := math.MaxInt32
		for j := 0; j < g.numVertices; j++ {
			if !final[j] && dist[j] < min {
				min = dist[j]
				idx = j
			}
		}
		if idx == -1 {
			break
		}
		final[idx] = true

		//检查所有邻接自pre的所有点，若其为final为false，这更新和path信息
		for _, edge := range g.GetEdges(idx) {
			//更新如果这个点没有确认，且距离到from的距离小于 from->v-> 这个点的距离 ，则更新这个点最小距离
			if !final[edge.dest] && edge.weight+dist[idx] < dist[edge.dest] {
				dist[edge.dest] = edge.weight + dist[idx]
				fromPath[edge.dest] = idx
			}
		}
	}
	return dist[to]
}
func main() {
	graph := NewGraph()

	// 添加顶点
	for i := 0; i < 6; i++ {
		graph.AddVertex(i)
	}
	// 添加边
	graph.AddEdge(0, 1, 3)
	graph.AddEdge(0, 2, 1)
	graph.AddEdge(0, 3, 2)
	graph.AddEdge(1, 4, 2)
	graph.AddEdge(1, 5, 3)
	graph.AddEdge(2, 4, 4)
	graph.AddEdge(3, 5, 9)
	minPath := graph.DijkStraMinPath(0, 5)
	fmt.Println(minPath)
}
