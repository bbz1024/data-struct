package insert

/*
每次将一个待排序的记录按其关键字大小插入到前面已经排好序的子序列中，直到全部记录插入完成。
特点：2轮循环之前的都是有序了的
适用于：链表，线性表
*/

func Sort(nums []int) {
	// 外循环：已排序区间为 [0, i-1]
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		// 内循环：将 base 插入到已排序区间 [0, i-1] 中的正确位置
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j] // 将 nums[j] 向右移动一位
			j--
		}
		nums[j+1] = base // 将 base 赋值到正确位置
	}
}

func Sort2(nums []int) {

	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i - 1
		for ; j >= 0 && temp < nums[j]; j-- {
			nums[j+1] = nums[j] // 向左调整
			//nums[j+1], nums[j] = nums[j], nums[j+1]
		}
		//调整完成进行插入到合适的位置
		nums[j+1] = temp
	}
}
