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
	var left, right *TreeNode
	stack := list.New()
	for root != nil || stack.Len() > 0 {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		root = stack.Back().Value.(*TreeNode)

	}
}

// @lc code=end
func main() {

}
