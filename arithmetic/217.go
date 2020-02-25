package arithmetic

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/2/22 18:09
 */
func ContainsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	var tempMap = map[int]bool{}
	for _, v := range nums {
		if tempMap[v] {
			return true
		}
		tempMap[v] = true
	}
	return false
}
