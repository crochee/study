package arithmetic

import "fmt"

/*
* @
* @Author:
* @Date: 2020/3/20 20:24
 */
// 计数排序  元素是整型，数据范围小（最大-最小的值）  时间复杂度o(n+m) 空间复杂度o(m)
func CountSort(nums []int) []int {
	var (
		max = nums[0]
		min = nums[0]
	)
	for _, v := range nums {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	d := max - min
	countArr := make([]int, d+1)
	for _, v := range nums {
		countArr[v-min]++
	}
	fmt.Println(countArr)
	for i := 1; i < len(countArr); i++ {
		countArr[i] += countArr[i-1]
	}
	sortArr := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		sortArr[countArr[nums[i]-min]-1] = nums[i]
		countArr[nums[i]-min]--
	}
	fmt.Println(countArr)
	return sortArr
}
