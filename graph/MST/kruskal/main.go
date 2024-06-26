package main

import "fmt"

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, size),
		rank:   make([]int, size),
	}
	for i := range uf.parent {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX == rootY {
		return
	}

	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
}

func main() {
	uf := NewUnionFind(5)

	uf.Union(0, 1)
	uf.Union(1, 2)
	uf.Union(1, 4)
	uf.Union(4, 3)
	//uf.Union(3, 4)

	fmt.Println("Find(0):", uf.Find(0))
	fmt.Println("Find(2):", uf.Find(2))
	fmt.Println("Find(2):", uf.Find(3))
	fmt.Println("Find(3):", uf.Find(4))
}
