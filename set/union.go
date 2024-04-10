package main

import "fmt"

func main() {
	union := InitUnion(10)
	union.Union(1, 2)
	union.Union(1, 3)
	union.Union(2, 4)
	union.Union(5, 6)

	fmt.Println(union.Find(6))
	fmt.Println(union.Find(4))
}

type Union struct {
	Set []int
}

func InitUnion(size int) *Union {
	u := &Union{
		Set: make([]int, size),
	}
	for i := 0; i < size; i++ {
		u.Set[i] = -1 // 指向自己
	}
	return u
}
func (u *Union) Find(v int) int {
	//找到父了
	if u.Set[v] == -1 {
		return v
	}
	return u.Find(u.Set[v])
}
func (u *Union) Union(v, v1 int) {
	if v == v1 {
		return
	}
	u.Set[u.Find(v1)] = u.Find(v)
}
