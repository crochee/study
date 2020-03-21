package arithmetic

/*
* @
* @Author:
* @Date: 2020/3/20 22:44
 */

func HeapSort(nums []int) {
	for i := (len(nums) - 2) / 2; i >= 0; i-- {
		downJust(nums, i, len(nums))
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		downJust(nums, 0, i)
	}
}

func downJust(nums []int, parentIndex, length int) {
	var (
		temp           = nums[parentIndex]
		leftChildIndex = 2*parentIndex + 1
	)
	for leftChildIndex < length {
		if leftChildIndex+1 < length && nums[leftChildIndex+1] > nums[leftChildIndex] {
			leftChildIndex++
		}
		if temp >= nums[leftChildIndex] {
			break
		}
		nums[parentIndex] = nums[leftChildIndex]
		parentIndex = leftChildIndex
		leftChildIndex = 2*leftChildIndex + 1
	}
	nums[parentIndex] = temp
}
