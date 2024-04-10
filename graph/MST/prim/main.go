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

// PrimMST Prim算法普里姆（Prim）算法计算最小生成树。
func (g *Graph) PrimMST(start int) int {
	// 如果图为空或者没有顶点，则返回 0
	if g.numVertices == 0 {
		return 0
	}

	// 初始化
	cost := 0
	lowCost := make([]int, g.numVertices)
	for i := range lowCost {
		lowCost[i] = math.MaxInt32
	}
	isJoin := make([]bool, g.numVertices)
	isJoin[start] = true
	// 更新起始顶点的权重
	edges := g.GetEdges(start)
	for _, edge := range edges {
		lowCost[edge.dest] = edge.weight
	}
	// 主循环
	for i := 0; i < g.numVertices-1; i++ {
		// 找到未加入最小生成树的顶点中权重最小的顶点
		min := math.MaxInt32
		low := -1
		for j := 0; j < g.numVertices; j++ {
			if !isJoin[j] && lowCost[j] < min {
				min = lowCost[j]
				low = j
			}
		}

		// 标记该顶点为已加入
		isJoin[low] = true
		cost += min

		// 更新最小权重数组
		for _, edge := range g.GetEdges(low) {
			if !isJoin[edge.dest] && edge.weight < lowCost[edge.dest] {
				lowCost[edge.dest] = edge.weight
			}
		}
	}

	return cost
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

	for i := 0; i < 6; i++ {
		fmt.Println(graph.PrimMST(i))
	}
}
