package main

import (
	"fmt"
	"math"
)

type MGraph struct {
	Node []int   // 顶点表
	Edge [][]int // 边
}

// AddEdges 添加边
func (m *MGraph) AddEdge(node1, node2 int, weight int) {
	node1 = m.Node[node1]
	node2 = m.Node[node2]
	//node1 -> node2
	m.Edge[node1][node2] = weight
}
func (m *MGraph) AddVertex(node1 int) {
	if node1 > len(m.Node)-1 && node1 < 0 {
		return
	}
	m.Node[node1] = node1
}

// GetEdge 获取度
func (m *MGraph) GetEdge(vertex int) []int {
	var edges []int
	for i := 0; i < len(m.Node); i++ {
		val := m.Edge[vertex][i]
		if val != math.MaxInt32 {
			edges = append(edges, m.Node[i])
		}
	}
	return edges
}

func Init(size int) *MGraph {
	graph := &MGraph{
		Node: make([]int, size),
		Edge: make([][]int, size),
	}
	for i := 0; i < size; i++ {
		graph.Edge[i] = make([]int, size)
		for j := 0; j < size; j++ {
			graph.Edge[i][j] = math.MaxInt32
		}
	}
	return graph
}
func (m *MGraph) Floyd(to, from int) int {
	l := len(m.Node)
	fromPath := make([][]int, l)
	for i := 0; i < l; i++ {
		fromPath[i] = make([]int, l)
		for j := 0; j < l; j++ {
			fromPath[i][j] = -1
		}
	}
	for k := 0; k < l; k++ { // 增加新的中转点
		for i := 0; i < l; i++ { // 遍历整个矩阵，i为行号，j为列号
			for j := 0; j < l; j++ {
				if m.Edge[i][j] > m.Edge[i][k]+m.Edge[k][j] { // 以vk为中转点的路径更短
					m.Edge[i][j] = m.Edge[i][k] + m.Edge[k][j] // 更新最短路径长度
					fromPath[i][j] = k                         //中转点
				}
			}
		}
	}
	//递归获取顶点经过的点。
	/*
		假如从0顶点到3顶点：
			fromPath[0][3]==2: 0 是经过2过来最短路径
			查找0-2 2-3 是否存在最短路径
			fromPath[0][2] == -1: 为-1就是已经是最短路径了。
			fromPath[2][3] == 1: 2->3 存在中转顶点1
			查找2-1 1-3 是否存在最短路径
			fromPath[2][1] == -1;
			fromPath[1][3] == -1;
			此时0顶点到3顶点其中需要经过
			0		2（是最短的）		3
			0 -> 2 不存在中转顶点
			2 -> 3 存在中转顶点
			2	1	3
			2-> 1 不存在中转顶点
			1-> 3 不存在中转顶点
			所以：
				pass ：0 -> 2 -> 1 -> 3
		fromPath:
			0 [-1 2 -1 2 3]
			1 [-1 -1 -1 -1 3]
			2 [-1 -1 -1 1 3]
			3 [-1 -1 -1 -1 -1]
			4 [-1 -1 -1 -1 -1]
	*/

	fmt.Println(fromPath)
	m.Pass(to, from, fromPath)

	//其中最短路径经过的顶点
	return m.Edge[to][from]
}

// Pass 其中最短路径经过的顶点
func (m *MGraph) Pass(from, to int, path [][]int) {
	if path[from][to] == -1 {
		return
	}
	//如果
	m.Pass(from, path[from][to], path)
	fmt.Print(m.Node[path[from][to]], " ")
	m.Pass(path[from][to], to, path)
}

func main() {
	graph := Init(5)
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(0, 2, 1)
	graph.AddEdge(0, 4, 10)
	graph.AddEdge(1, 3, 1)
	graph.AddEdge(1, 4, 5)
	graph.AddEdge(2, 1, 1)
	graph.AddEdge(2, 4, 7)
	graph.AddEdge(3, 4, 1)

	res := graph.Floyd(0, 3)
	fmt.Println()
	fmt.Println(res)
}
