package arithmetic

/**
* @ Description:旋转数组
* @Author: licongfu@ringle.com
* @Date: 2020/2/22 14:58
 */
// 1.使用环状替换
func Rotate1(nums []int, k int) {
	k = k % len(nums)
	var count int
	for start := 0; count < len(nums); start++ {
		current := start
		prev := nums[start]
		next := (current + k) % len(nums)
		temp := nums[next]
		nums[next] = prev
		prev = temp
		current = next
		count++
		for start != current {
			next = (current + k) % len(nums)
			temp = nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
		}
	}
}

// 2.3次翻转

func Rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1) // 反转所有数据
	reverse(nums, 0, k-1)         // 反转0->k-1
	reverse(nums, k, len(nums)-1) // 反转k->n-1
}

func reverse(nums []int, start, end int) { // 反转的方法，前后对调
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
