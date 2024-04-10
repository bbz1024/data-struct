package main

import "fmt"

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

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func newNode(key int) *Node {
	return &Node{
		Key:    key,
		Left:   nil,
		Right:  nil,
		Height: 1,
	}
}
func rightRotate(y *Node) *Node {
	//
	/*
						6
				5				12
			2		3
		1
	*/
	x := y.Left   //5
	t2 := x.Right //3

	x.Right = y //12
	y.Left = t2

	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

func leftRotate(x *Node) *Node {
	y := x.Right
	t2 := y.Left

	y.Left = x
	x.Right = t2

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

func (tree *AVLTree) insert(node *Node, key int) *Node {
	if node == nil {
		return newNode(key)
	}

	if key < node.Key {
		node.Left = tree.insert(node.Left, key)
	} else if key > node.Key {
		node.Right = tree.insert(node.Right, key)
	} else {
		return node
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))

	balance := getBalance(node)

	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}

	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}

	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func (tree *AVLTree) Insert(key int) {
	tree.Root = tree.insert(tree.Root, key)
}

func preOrder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.Key)
		preOrder(root.Left)
		preOrder(root.Right)
	}
}

func main() {
	tree := AVLTree{}

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(25)

	fmt.Println("Preorder traversal of the constructed AVL tree is:")
	preOrder(tree.Root)
}
