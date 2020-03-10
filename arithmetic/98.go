package arithmetic

import "math"

/**
* @ Description:
* @Author:
* @Date: 2020/3/2 14:18
 */

type TreeNode98 struct {
	Val   int
	Left  *TreeNode98
	Right *TreeNode98
}

func isValidBST(root *TreeNode98) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode98, lower, upper int) bool {
	if node == nil {
		return true
	}
	if node.Val <= lower || node.Val >= upper {
		return false
	}
	return helper(node.Right, node.Val, upper) && helper(node.Left, lower, node.Val)
}
