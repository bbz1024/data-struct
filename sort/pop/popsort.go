package pop

/*
从前往后凉凉比较元素的值，并且交换他们，直到
https://visualgo.net/zh/sorting
https://www.hello-algo.com/chapter_sorting/bubble_sort/
*/

func Sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		modify := true
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				modify = false
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
		//优化部分，如果不存在交换值就说明已经改变了
		if !modify {
			break
		}
	}
}
