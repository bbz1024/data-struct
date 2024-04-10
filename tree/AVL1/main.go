package main

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Height int
}

type AVLTree struct {
	Root *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newNode(key int) *Node {
	return &Node{
		Key:    key,
		Left:   nil,
		Right:  nil,
		Height: 1,
	}
}

func main() {
	tree := AVLTree{}

}
