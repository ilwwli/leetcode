/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(postorder) != len(inorder) {
		return nil
	}
	n := len(postorder)
	ind := 0
	for ; ind < len(inorder); ind++ {
		if postorder[n-1] == inorder[ind] {
			break
		}
	}
	return &TreeNode{
		Val:   inorder[ind],
		Left:  buildTree(inorder[:ind], postorder[:ind]),
		Right: buildTree(inorder[ind+1:], postorder[ind:n-1]),
	}
}

// @lc code=end
