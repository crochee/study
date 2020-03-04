package arithmetic

/**
* @ Description:
* @Author: licongfu@ringle.com
* @Date: 2020/3/4 17:18
 */

type TreeNode102 struct {
	Val   int
	Left  *TreeNode102
	Right *TreeNode102
}

func levelOrder(root *TreeNode102) [][]int {
	levels := make([][]int, 0)
	if root == nil {
		return levels
	}
	return helper102(root, levels, 0)
}

func helper102(node *TreeNode102, levels [][]int, level int) [][]int {
	if len(levels) == level {
		levels = append(levels, []int{})
	}
	levels[level] = append(levels[level], node.Val)
	if node.Left != nil {
		levels = helper102(node.Left, levels, level+1)
	}
	if node.Right != nil {
		levels = helper102(node.Right, levels, level+1)
	}
	return levels
}
