package arithmetic

/*
* @
* @Author:冒泡排序系列
* @Date: 2020/3/20 17:08
 */

// 普通冒泡排序 时间复杂度o(n*n)
func BubbleSort1(nums []int) {
	for i := 0; i < len(nums); i++ { // 内层循环的次数
		for j := 0; j < len(nums)-i-1; j++ { // 左边大于右边交换
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// 冒泡排序优化版1
func BubbleSort2(nums []int) {
	for i := 0; i < len(nums); i++ { // 内层循环的次数
		var flag bool // 加标志看一轮里是否有交换，来判断是否有序
		for j := 0; j < len(nums)-i-1; j++ { // 左边大于右边交换
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// 冒泡排序优化版2
func BubbleSort3(nums []int) {
	var (
		lastExchangeIndex = 0             // 记录一轮最后一次交换的位置
		sortBorder        = len(nums) - 1 // 界定有无序的边界
	)
	for i := 0; i < len(nums); i++ { // 内层循环的次数
		var flag bool
		for j := 0; j < sortBorder; j++ {
			if nums[j] > nums[j+1] { // 左边大于右边交换
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
				lastExchangeIndex = j
			}
		}
		sortBorder = lastExchangeIndex
		if !flag {
			break
		}
	}
}

// 鸡尾酒排序 时间复杂度o(n*n)
func BubbleSort4(nums []int) {
	for i := 0; i < len(nums)/2; i++ { // 内层循环的次数减半
		var flag bool
		for j := i; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] { // 左边大于右边交换
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
		flag = false
		for j := len(nums) - i - 1; j > i; j-- {
			if nums[j] < nums[j-1] { // 左边大于右边交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// 鸡尾酒排序优化版1
func BubbleSort5(nums []int) {
	var (
		leftlastExchangeIndex  = 0
		rightlastExchangeIndex = 0
		leftsortBorder         = len(nums) - 1
		rightsortBorder        = 0
	)
	for i := 0; i < len(nums)/2; i++ { // 内层循环的次数减半
		var flag bool
		for j := rightsortBorder; j < leftsortBorder; j++ {
			if nums[j] > nums[j+1] { // 左边大于右边交换
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
				leftlastExchangeIndex = j
			}
		}
		leftsortBorder = leftlastExchangeIndex
		if !flag {
			break
		}
		flag = false
		for j := leftsortBorder; j > rightsortBorder; j-- {
			if nums[j] < nums[j-1] { // 左边大于右边交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
				flag = true
				rightlastExchangeIndex = j
			}
		}
		rightsortBorder = rightlastExchangeIndex
		if !flag {
			break
		}
	}
}
