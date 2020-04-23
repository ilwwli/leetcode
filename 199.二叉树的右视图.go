package main

import (
	"container/list"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	sentinal := &TreeNode{}
	l := list.New()
	l.PushBack(root)
	l.PushBack(sentinal)
	lastval := 0
	for l.Len() > 0 {
		node := l.Front().Value.(*TreeNode)
		l.Remove(l.Front())
		if node == sentinal {
			ans = append(ans, lastval)
			if l.Len() > 0 {
				l.PushBack(sentinal)
			}
			continue
		}
		lastval = node.Val
		if node.Left != nil {
			l.PushBack(node.Left)
		}
		if node.Right != nil {
			l.PushBack(node.Right)
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

func main() {
	root := &TreeNode{Val: 1}

	fmt.Println(rightSideView(root))
}
