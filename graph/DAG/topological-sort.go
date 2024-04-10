package main

import (
	"fmt"
)

type MGraph struct {
	Node []int   // 顶点表
	Edge [][]int // 边
}

// AddEdge  添加边
func (m *MGraph) AddEdge(node1, node2 int) {
	node1 = m.Node[node1]
	node2 = m.Node[node2]
	//node1 -> node2
	m.Edge[node1][node2] = 1
}

// AddVertex 添加顶点
func (m *MGraph) AddVertex(node1 int) {
	if node1 > len(m.Node)-1 && node1 < 0 {
		return
	}
	m.Node[node1] = node1
}

// GetOutEdge 获取出度
func (m *MGraph) GetOutEdge(vertex int) []int {
	var edges []int
	for i := 0; i < len(m.Node); i++ {
		val := m.Edge[vertex][i]
		if val == 1 {
			edges = append(edges, m.Node[i])
		}
	}
	return edges
}

// GetInEdge 获取入度
func (m *MGraph) GetInEdge(vertex int) []int {
	var edges []int
	for i := 0; i < len(m.Node); i++ {
		val := m.Edge[i][vertex]
		if val == 1 {
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
	}
	return graph
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
func (m *MGraph) AOV() {
	var stack []int
	//1. 找到不存在入度的节点，加入栈
	l := len(m.Node)
	inDegree := make([]int, l)

	// TODO ----------- Start ----------- 以下代码存在性能问题，及并没有更好的利用cpu的预读策略
	for i := 0; i < l; i++ {
		sum := 0
		for j := 0; j < l; j++ {
			sum += m.Edge[j][i] // Issues
		}
		if sum == 0 {
			stack = append(stack, i)
		}
		//记录所有顶点的入度数量
		inDegree[i] = sum
		//
	}
	// ----------- END -----------

	printArr := make([]int, l)
	var count int
	for len(stack) != 0 {
		//出栈一个元素
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		printArr[count] = val // 输出到顶点
		count++
		for _, i := range m.GetOutEdge(val) {
			inDegree[i]--
			// 将顶点val的出度-1，并且将度为0的顶点加入栈
			if inDegree[i] == 0 {
				stack = append(stack, i)
			}
		}
	}
	if count != l {
		fmt.Println("排序失败，是一个有环图")
	} else {
		for _, p := range printArr {
			fmt.Print(aov[p])
			fmt.Print(" -> ")
		}
	}
}

// DFSTraverse 逆拓扑排序
func (m *MGraph) DFSTraverse() {
	visited := make([]bool, len(m.Node))
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			m.DFSReverse(i, visited)
		}
	}
}
func (m *MGraph) DFSReverse(i int, visited []bool) {
	visited[i] = true
	edges := m.GetOutEdge(i)
	for _, edge := range edges {
		if !visited[edge] {
			m.DFSReverse(edge, visited)
		}
	}
	fmt.Println(aov[i])

}

//func (m *MGraph) DFSSort() {
//	visited := make([]bool, len(m.Node))
//	m.DFS(4, visited)
//}
//
//func (m *MGraph) DFS(i int, visited []bool) {
//	visited[i] = true
//	for _, v := range m.GetInEdge(i) {
//		if !visited[v] {
//			m.DFS(v, visited)
//		}
//	}
//	fmt.Println(aov[i])
//	for _, v := range m.GetOutEdge(i) {
//		if !visited[v] {
//			m.DFS(v, visited)
//		}
//	}
//}

var aov = [...]string{"准备厨具", "打鸡蛋", "下锅炒", "吃", "买菜", "洗西蕃茄", "切番茄"}

func main() {
	graph := Init(7)
	graph.AddVertex(0) // 准备厨具
	graph.AddVertex(1) // 打鸡蛋
	graph.AddVertex(2) // 下锅炒
	graph.AddVertex(3) // 吃
	graph.AddVertex(4) // 买菜
	graph.AddVertex(5) // 洗西蕃茄
	graph.AddVertex(6) // 切番茄
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 6)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(4, 1)
	graph.AddEdge(4, 5)
	graph.AddEdge(5, 6)
	graph.AddEdge(6, 2)

	//graph.AOV()
	//graph.DFSTraverse()
	//graph.DFSSort()
	//fmt.Println(graph.GetOutEdge(6))
	//fmt.Println(graph.GetInEdge(6))
}
