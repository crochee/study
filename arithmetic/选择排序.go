package arithmetic

/*
* @
* @Author:
* @Date: 2020/3/20 22:20
 */

// 选择排序
func SelectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] { // 剩余元素跟当前元素比较
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
}
