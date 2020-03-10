package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/22 18:38
 */

func SingleNumber(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	var temp int
	for _, v := range nums {
		temp ^= v
	}
	return temp
}
