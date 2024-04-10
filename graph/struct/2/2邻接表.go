package main

import "fmt"

// EdgeNode 边表结点
type EdgeNode struct {
	index int
	next  *EdgeNode // 指向下一个邻接点
}

// VertexNode 顶点表结点
type VertexNode struct {
	val   int       // 顶点域
	first *EdgeNode // 边表头指针
	idx   int
}

type Graph []*VertexNode

func (g *Graph) InsertNode(node *VertexNode) {
	node.idx = len(*g)
	node.first = &EdgeNode{
		index: -1,
	}
	*g = append(*g, node)
}

// GetEdge 获取某个节点的边
func (g *Graph) GetEdge(node VertexNode) {
	temp := node.first.next
	for temp != nil {
		fmt.Println((*g)[temp.index].val)
		temp = temp.next
	}

}
func (g *Graph) InsertEdge(node1, node2 *VertexNode) {
	if node1 == nil || node2 == nil || node1 == node2 {
		return
	}
	edgeNode1 := &EdgeNode{
		index: node2.idx,
		next:  node1.first.next,
	}
	node1.first.next = edgeNode1
	edgeNode2 := &EdgeNode{
		index: node1.idx,
		next:  node2.first.next,
	}
	node2.first.next = edgeNode2
}

func (g *Graph) DeleteNode(node *VertexNode) {
	if node == nil || node.idx >= len(*g) {
		return
	}

	//123456
	//456,56
	// 从切片中移除节点
	copy((*g)[node.idx:], (*g)[node.idx+1:])
	*g = (*g)[:len(*g)-1]

	// 调整其他节点的索引
	for i := node.idx; i < len(*g); i++ {
		(*g)[i].idx = i
	}
	fmt.Println(*g)
	// 查看其他节点是否存在指向该 node 的边，并删除它们
	for i := 0; i < len(*g); i++ {
		if nd := (*g)[i]; nd != nil {
			temp := nd.first.next
			for temp != nil {
				if temp.next != nil && temp.next.index == node.idx {
					temp.next = temp.next.next
					break
				}
				temp = temp.next
			}
		}
	}
}

func main() {
	g := Graph{}
	node1 := VertexNode{
		val: 1,
	}
	node2 := VertexNode{
		val: 2,
	}
	node3 := VertexNode{
		val: 3,
	}
	node4 := VertexNode{
		val: 4,
	}
	g.InsertNode(&node1)
	g.InsertNode(&node2)
	g.InsertEdge(&node1, &node2)
	g.InsertNode(&node3)
	g.InsertEdge(&node2, &node3)
	//g.GetEdge(node2)
	g.InsertNode(&node4)
	//g.GetEdge(node4)
	g.InsertEdge(&node1, &node4)
	g.InsertEdge(&node2, &node4)
	g.InsertEdge(&node3, &node4)
	//g.InsertEdge(&node3, &node4)
	//g.GetEdge(node4)
	g.DeleteNode(&node1)
	g.GetEdge(node4)
}

/*
// 创建一个切片
slice := []int{1, 2, 3, 4, 5}

// 要删除的元素的索引
index := 2

// 使用 copy 函数将要删除的元素之后的元素向前移动一位
copy(slice[index:], slice[index+1:])

[3,4,5] [4,5] [3,4,5,4,5]

[]
// 将切片的长度减一，删除最后一个元素
slice = slice[:len(slice)-1]
*/
