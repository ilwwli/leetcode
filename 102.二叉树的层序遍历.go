import "container/list"

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := list.New()
	dummy := &TreeNode{}
	queue.PushBack(root)
	queue.PushBack(dummy)
	var tmp []int
	for queue.Len() > 0 {
		ele := queue.Front()
		queue.Remove(ele)
		node := ele.Value.(*TreeNode)
		if node == dummy {
			res = append(res, tmp)
			tmp = nil
			if queue.Len() > 0 {
				queue.PushBack(dummy)
			}
			continue
		}
		tmp = append(tmp, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return res
}

// @lc code=end
