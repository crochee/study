package arithmetic

/*
* @
* @Author:
* @Date: 2020/3/20 18:45
 */
// 快排递归版本 n*log(n)
func QuickSort(nums []int, start, end int) {
	if start >= end { // 结束条件
		return
	}
	// 获取基准元素位置
	pivotIndex := partition1(nums, start, end)
	QuickSort(nums, start, pivotIndex-1) // 左边
	QuickSort(nums, pivotIndex+1, end)   // 右边
}

// 分治  双边循环法
func partition(nums []int, start, end int) int {
	var (
		pivot = nums[start] // 选择基准元素，此出用第一个元素，也可以随机选择
		left  = start
		right = end
	)
	for left != right { // 左右两边
		for left < right && nums[right] > pivot {
			right--
		}
		for left < right && nums[left] <= pivot {
			left++
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	// 交换基准元素
	nums[start], nums[left] = nums[left], pivot
	return left
}

// 分治 单边循环法
func partition1(nums []int, start, end int) int {
	var (
		pivot = nums[start] // 选择基准元素，此出用第一个元素，也可以随机选择
		mark  = start
	)
	for i := start + 1; i <= end; i++ { // 左右两边
		if nums[i] < pivot {
			mark++
			nums[mark], nums[i] = nums[i], nums[mark]
		}
	}
	// 交换基准元素
	nums[start], nums[mark] = nums[mark], pivot
	return mark
}
