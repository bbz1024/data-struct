package main

func main() {

}

type Union2 struct {
	Set []int
}

func InitUnion2(size int) *Union2 {
	u := &Union2{Set: make([]int, size)}
	for i := 0; i < size; i++ {
		u.Set[i] = -i
	}
	return u
}
func (u *Union2) Find(v int) int {
	for u.Set[v] >= 0 {
		v = u.Set[v]
	}
	return v
}
func (u *Union2) Union(root1, root2 int) {
	if root1 == root2 {
		return
	}
	u.Set[root2] = root1 // 标记v1为父
}
func (u *Union2) UnionPro(root1, root2 int) {
	if root1 == root2 {
		return
	}
	//root2结点数更少
	if u.Set[root2] > u.Set[root1] {
		//累加结点总数
		u.Set[root1] += u.Set[root2]
		//小树合并到大树
		u.Set[root2] = root1
	} else {
		u.Set[root2] += u.Set[root1]
		u.Set[root1] = root2
	}
}
