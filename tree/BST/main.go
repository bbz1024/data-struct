package main

import "fmt"

type BstNode struct {
	Val         int
	Left, Right *BstNode
}

type BST struct {
	root *BstNode
}

func NewBst() *BST {
	return &BST{}
}
func (b *BST) Insert(val int) {
	if b.root == nil {
		b.root = &BstNode{
			Val: val,
		}
		return
	}
	b.root.insert(val)
}
func (b *BST) Delete(val int) {
	if b.root == nil {
		return
	}
	b.root.delete(val)
}
func (b *BstNode) delete(val int) {
	temp := b
	var pre *BstNode // 保存父节点
	for temp != nil {
		if temp.Val == val {
			break
		}
		pre = temp
		if temp.Val < val {
			temp = temp.Right
		} else if temp.Val > val {
			temp = temp.Left
		}
	}
	if temp == nil {
		return
	}
	//此时的temp就是删除的节点
	//1. 为叶子节点，则直接修改父的其中一个孩子指向nil
	if temp.Left == temp.Right {
		if pre.Left == temp {
			pre.Left = nil
		} else {
			pre.Right = nil
		}
		return
	}
	//2. 存在一个孩子节点或者2个节点
	if temp.Left == nil || temp.Right == nil {
		var replacement *BstNode
		if temp.Left != nil {
			replacement = temp.Left
		} else {
			replacement = temp.Right
		}
		if pre.Left == temp {
			pre.Left = replacement
		} else {
			pre.Right = replacement
		}
	} else {
		//存在2个节点（度为2）
		//前置或后置进行补充
		//前置：左孩子的最右节点
		//后置：右孩子的最左节点
		//1.前置替代
		left := temp.Left
		for left.Right != nil {
			left = left.Right
		}
		//递归删除这个节点
		b.delete(left.Val)
		temp.Val = left.Val
	}
	return
}
func (b *BstNode) insert(val int) *BstNode {
	if b == nil {
		return &BstNode{
			Val: val,
		}
	}
	if b.Val > val {
		b.Left = b.Left.insert(val)
	} else if b.Val < val {
		b.Right = b.Right.insert(val)
	}
	return b
}

func main() {
	st := NewBst()
	st.Insert(3)
	st.Insert(2)
	st.Insert(5)
	st.Insert(4)
	st.Delete(3)
	fmt.Println(st.root)
}
