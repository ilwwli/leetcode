package main

import (
	"container/list"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=958 lang=golang
 *
 * [958] 二叉树的完全性检验
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
func isCompleteTree(root *TreeNode) bool {
	queue := list.New()
	queue.PushBack(root)
	flag := false
	for ; queue.Front() != nil; queue.Remove(queue.Front()) {
		v := queue.Front().Value.(*TreeNode)
		if v == nil {
			flag = true
			continue
		}
		if flag {
			return false
		}
		queue.PushBack(v.Left)
		queue.PushBack(v.Right)
	}
	return true
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	left := &TreeNode{}
	root := &TreeNode{Left: left}

	fmt.Println(isCompleteTree(root))
}
