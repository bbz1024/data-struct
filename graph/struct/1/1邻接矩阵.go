package main

import (
	"fmt"
	"strings"
)

const MaxNode = 6

type MGraph struct {
	Node [MaxNode]byte         // 顶点表
	Edge [MaxNode][MaxNode]int // 边
}

var visited = [MaxNode]bool{}

// AddVertex 添加边
func (m *MGraph) AddVertex(node1, node2 byte) {
	node1 = m.Node[node1-1]
	node2 = m.Node[node2-1]
	// 无向图
	m.Edge[node1][node2], m.Edge[node2][node1] = 1, 1
}
func (m *MGraph) DeleteNode(node byte) {
	deleteIndex := int(node - 1)

	// 清空待删除节点所在行和列
	for i := range m.Edge[deleteIndex] {
		m.Edge[deleteIndex][i] = 0
		m.Edge[i][deleteIndex] = 0
	}
	// 遍历其他行，删除指向待删除节点的边
	for i := 0; i < MaxNode; i++ {
		if i != deleteIndex {
			m.Edge[i][deleteIndex] = 0
		}
	}
}

// BFS 广度优先遍历，类似与树的层级遍历
func (m *MGraph) BFS(start int) {
	var queue []int
	fmt.Println(m.Node[start-1]) // 访问
	visited[start-1] = true
	queue = append(queue, start)
	for len(queue) != 0 {
		//出队
		val := queue[0]
		queue = queue[1:] // 入队所有的出边
		// 访问
		for i := 0; i < MaxNode; i++ {
			isEdge := m.Edge[val][i]
			if isEdge != -1 && !visited[m.Node[i]] {
				visited[m.Node[i]] = true
				fmt.Println(m.Node[i])
				queue = append(queue, int(m.Node[i]))
			}
		}
	}
}

// BFS2 针对非连通图，广度优先遍历，
func (m *MGraph) BFS2() {
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			m.BFS(i + 1)
		}
	}
}

// GetEdge 获取度
func (m *MGraph) GetEdge(vertex int) []int {
	var edges []int
	for i := 0; i < MaxNode; i++ {
		val := m.Edge[vertex][i]
		if val != -1 {
			edges = append(edges, int(m.Node[i]))
		}
	}
	return edges
}

// DFS 类似与树的先根遍历（先序、中序、后序），深度优先
func (m *MGraph) DFS(start int) {
	//进行访问
	fmt.Println(m.Node[start])
	//标记为访问
	visited[start] = true
	//所有的出度
	edges := m.GetEdge(start)
	for _, edge := range edges {
		if !visited[edge] {
			//访问
			m.DFS(edge)
		}
	}
}

func (m *MGraph) DFS2() {
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			m.DFS(i)
		}
	}
}

func main() {

	//云南能投
	m := MGraph{
		Node: [MaxNode]byte{0, 1, 2, 3, 4, 5},
	}
	//插入边
	m.AddVertex(1, 2)
	m.AddVertex(1, 3)
	m.AddVertex(2, 3)
	m.AddVertex(2, 4)
	m.AddVertex(3, 5)

	fmt.Println(strings.Repeat("-", 50))
	//	删除一个节点
	//m.DeleteNode(1)
	//m.BFS2()
	fmt.Println(m.Node)
	//m.BFS(1)
	//m.DFS(3)
	m.DFS2()
}
