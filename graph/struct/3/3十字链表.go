package main

import "fmt"

type ArcNode struct {
	TailVex int      // 尾节点索引
	HeadVex int      // 头节点索引
	HLink   *ArcNode // 指向具有相同头节点的下一条边
	TLink   *ArcNode // 指向具有相同尾节点的下一条边
}

type VNode struct {
	Data     int      // 顶点数据
	Index    int      // 在图数组中的索引位置
	FirstIn  *ArcNode // 第一条入边，即指向该顶点的边链表头指针
	FirstOut *ArcNode // 第一条出边，即从该顶点出发的边链表头指针
}

// XLinkGraph 定义一个十字链表有向图类型
type XLinkGraph struct {
	Vertices      []*VNode // 存放所有顶点的数组
	NumOfVertices int      // 图中顶点数量
}

// AddVertex 添加顶点到图中
func (g *XLinkGraph) AddVertex(data int) {
	newNode := &VNode{
		Data:     data,
		Index:    g.NumOfVertices,
		FirstIn:  nil,
		FirstOut: nil,
	}
	g.Vertices = append(g.Vertices, newNode)
	g.NumOfVertices++
}

// AddEdge 添加有向边到图中，假设 tail 和 head 分别是边的尾节点和头节点的索引
func (g *XLinkGraph) AddEdge(tail, head int) {
	if tail < 0 || tail >= g.NumOfVertices || head < 0 || head >= g.NumOfVertices {
		fmt.Println("索引越界")
		return
	}
	if tail == head {
		fmt.Println("不能添加自环")
		return
	}

	tailNode := g.Vertices[tail]
	headNode := g.Vertices[head]

	newArc := &ArcNode{
		TailVex: tail,
		HeadVex: head,
	}

	// 将新边添加到头节点的入边链表 <-
	newArc.TLink = headNode.FirstIn
	headNode.FirstIn = newArc

	// 将新边添加到尾节点的出边链表 ->
	newArc.HLink = tailNode.FirstOut
	tailNode.FirstOut = newArc
}

func main() {
	xlink := XLinkGraph{}
	xlink.AddVertex(1)  // index: 0 A
	xlink.AddVertex(2)  // index: 1 B
	xlink.AddVertex(3)  // index: 2 C
	xlink.AddVertex(4)  // index: 3 D
	xlink.AddEdge(0, 1) // A -> B
	xlink.AddEdge(0, 2) // A -> C
	xlink.AddEdge(2, 0) // C -> A
	xlink.AddEdge(2, 3) // C -> D
	xlink.AddEdge(3, 3) // D -> C
	xlink.AddEdge(3, 1) // D -> B
	xlink.AddEdge(3, 0) // D -> A
	fmt.Println(xlink)

}
