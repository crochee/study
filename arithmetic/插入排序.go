package arithmetic

/*
* @
* @Author:
* @Date: 2020/3/20 22:08
 */

func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}
