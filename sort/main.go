package main

import (
	"data-struct/sort/selects"
	"fmt"
)

func main() {
	arr := []int{2, 1, 3, 5, 4, 6, 7, 9, 8}
	//insert.Sort2(arr)
	//pop.Sort(arr)
	//pop.Sort(arr)
	selects.Sort(arr)
	fmt.Println(arr)

}
