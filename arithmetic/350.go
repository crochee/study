package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/22 19:32
 */
func Intersect(nums1 []int, nums2 []int) []int {
	var temp1 = map[int]int{}
	for _, v := range nums1 {
		temp1[v]++
	}
	var arr = make([]int, 0)
	for _, v := range nums2 {
		if num, ok := temp1[v]; ok && num > 0 {
			arr = append(arr, v)
			temp1[v]--
		}
	}
	return arr
}
