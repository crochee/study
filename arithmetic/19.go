package arithmetic

/**
* @ Description:
* @Author:
* @Date: 2020/2/24 19:35
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	this := new(ListNode)
	this.Next = head
	first, second := this, this
	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return this.Next
}
