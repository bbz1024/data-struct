package selects

/*
每一趟在待排序元素中选取关键字最小的元素加入到有序子序列。
*/

func Sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ { // 一共进行n-1趟
		min := i                        // 记录最小元素位置
		for j := i; j < len(arr); j++ { // 在 arr[i,n-1]中选取最小的元素
			if arr[min] > arr[j] {
				//更新最小元素位置
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}
