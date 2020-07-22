// Copyright (c) Huawei Technologies Co., Ltd. 2020-2020. All rights reserved.
// Description:
// Author: l30002214
// Create: 2020/7/22

// Package arithmetic 
package arithmetic

import "testing"

func TestPreOrder(t *testing.T) {
	root := &BinaryTreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	root.Left = &BinaryTreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	root.Right = &BinaryTreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	temp1 := root.Left
	temp1.Left = &BinaryTreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	temp1.Right = &BinaryTreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	temp2 := root.Right
	temp2.Right = &BinaryTreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	temp1 = temp1.Right
	temp1.Left = &BinaryTreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	temp1.Right = &BinaryTreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	result := make([]int, 0)
	t.Log(PreOrder(root, result))
	result1 := make([]int, 0)
	t.Log(InOrder(root, result1))
	result2 := make([]int, 0)
	t.Log(AfterOrder(root, result2))
}
