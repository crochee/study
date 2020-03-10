package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/23 16:32
 */
func MoveZeroes(nums []int) {
	var j int
	for i, v := range nums {
		if v != 0 {
			if i != j {
				nums[j], nums[i] = nums[i], 0
			}
			j++
		}
	}
}
