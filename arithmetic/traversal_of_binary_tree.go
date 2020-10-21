package arithmetic

//二叉树的遍历:
//深度优先遍历:	前序 => 根节点--->左子树--->右子树
//				中序 => 左子树--->根节点--->右子树
//				后序 => 左子树--->右子树--->根节点
//广度优先遍历:	层次遍历=>按层遍历即可
type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

//前序 => 根节点--->左子树--->右子树
func PreOrder(root *BinaryTreeNode, result []int) []int {
	if root != nil {
		result = append(result, root.Val)
		result = PreOrder(root.Left, result)
		result = PreOrder(root.Right, result)
	}
	return result
}

//中序 = > 左子树--->根节点--->右子树
func InOrder(root *BinaryTreeNode, result []int) []int {
	if root != nil {
		result = InOrder(root.Left, result)
		result = append(result, root.Val)
		result = InOrder(root.Right, result)
	}
	return result
}

//中序 = > 左子树--->根节点--->右子树
func AfterOrder(root *BinaryTreeNode, result []int) []int {
	if root != nil {
		result = AfterOrder(root.Left, result)
		result = AfterOrder(root.Right, result)
		result = append(result, root.Val)
	}
	return result
}
