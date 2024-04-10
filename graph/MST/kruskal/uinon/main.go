package main

import "fmt"

/*
并查集：https://juejin.cn/post/7152763243868454920
*/
func main() {
	union := Init(8)
	union.Union(0, 1)
	union.Union(2, 1)
	union.Union(3, 1)
	union.Union(4, 3)
	union.Union(7, 6)
	fmt.Println(union.Find(6))

}

type UnionFind struct {
	parent []int
}

func Init(size int) *UnionFind {
	union := &UnionFind{
		parent: make([]int, size),
	}
	for i := 0; i < size; i++ {
		union.parent[i] = i // 先是指向自己
	}
	return union
}
func (u *UnionFind) Find(i int) int {
	if u.parent[i] == i {
		return i
	}
	return u.Find(u.parent[i])

}
func (u *UnionFind) Union(i, j int) {
	u.parent[u.parent[i]] = u.parent[j]
}
