package main

/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return []*TreeNode{}
	}
	return generateTreesRange(1, n+1)
}

func generateTreesRange(start, end int) []*TreeNode {
	if start >= end {
		return []*TreeNode{nil}
	}
	ans := make([]*TreeNode, 0)
	for i := start; i < end; i++ {
		leftNodes := generateTreesRange(start, i)
		rightNodes := generateTreesRange(i+1, end)
		for _, left := range leftNodes {
			for _, right := range rightNodes {
				tmp := &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				}
				ans = append(ans, tmp)
			}
		}
	}
	return ans
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
