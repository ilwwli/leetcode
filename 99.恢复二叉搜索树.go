package main

import "container/list"

/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverTree(root *TreeNode) {

	var left, right, last *TreeNode
	stack := list.New()
	for root != nil || stack.Len() > 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		root = stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())
		if last != nil {
			if last.Val > root.Val {
				right = root
				if left == nil {
					left = last
				}
			}
		}
		last = root
		root = root.Right
	}
	left.Val, right.Val = right.Val, left.Val
	return
}

// @lc code=end
func main() {

}
