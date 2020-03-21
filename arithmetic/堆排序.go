package arithmetic

/*
* @
* @Author:
* @Date: 2020/3/20 22:44
 */
// 从小到大的堆排序，需要最大堆
func HeapSort(nums []int) {
	// 构成最大堆.
	for i := len(nums)/2 - 1; i >= 0; i-- { // 只有n/2-1 ~ 0 有叶子节点
		downJust1(nums, i, len(nums))
	}
	// 需要将堆顶的最大的数据调整到尾部
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i] // 首尾交换
		downJust1(nums, 0, i)               // 此出i限定调整的范围,首尾交换后，将最后的元素屏蔽
	}
}

// 模拟二叉堆 广度优先遍历
func downJust(nums []int, parentIndex, length int) { // 作用是将父节点和两子节点比较，最大的值和父节点交换
	var (
		left  int // 左子节点的Index
		right int // 右子节点的Index
		temp  int // 记录选中子节点的Index
	)
	for 2*parentIndex+1 < length { // 判断是否有左子节点
		left = 2*parentIndex + 1 // 记录左子节点
		right = left + 1         // 记录右子节点
		temp = left              // 初始选中左子节点
		if right < length && nums[right] > nums[left] { // 如果有右子节点且右子节点大于左子节点，则记录右子节点index
			temp = right
		}
		if nums[parentIndex] < nums[temp] { // 如果父节点小于子节点的最大者，则交换并且父节点定位到交换的子节点
			nums[parentIndex], nums[temp], parentIndex = nums[temp], nums[parentIndex], temp
		} else {
			break
		}
	}
}

// 删除多余变量的改进
func downJust1(nums []int, parentIndex, length int) { // 作用是将父节点和两子节点比较，最大的值和父节点交换
	var left int // 左子节点的Index
	for 2*parentIndex+1 < length { // 判断是否有左子节点,循环的目的是需要保证子节点的子节点也需要是这样的规则
		left = 2*parentIndex + 1 // 记录左子节点
		if left+1 < length && nums[left+1] > nums[left] { // 如果有右子节点且右子节点大于左子节点，则记录右子节点index
			left++
		}
		if nums[parentIndex] >= nums[left] {
			break
		}
		// 如果父节点小于子节点的最大者，则交换并且父节点定位到交换的子节点
		nums[parentIndex], nums[left], parentIndex = nums[left], nums[parentIndex], left
	}
}
