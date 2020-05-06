import "container/list"

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层次遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := list.New()
	dummy := &TreeNode{}
	queue.PushBack(root)
	queue.PushBack(dummy)
	var tmp []int
	flag := true // true means left->right, false means reverse
	for queue.Len() > 0 {
		ele := queue.Front()
		if !flag {
			ele = queue.Back()
		}

		queue.Remove(ele)
		node := ele.Value.(*TreeNode)
		if node == dummy {
			res = append(res, tmp)
			tmp = nil
			if queue.Len() > 0 {
				if flag {
					queue.PushFront(dummy)
				} else {
					queue.PushBack(dummy)
				}
			}
			flag = !flag
			continue
		}
		tmp = append(tmp, node.Val)
		if flag {
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		} else {
			if node.Right != nil {
				queue.PushFront(node.Right)
			}
			if node.Left != nil {
				queue.PushFront(node.Left)
			}
		}

	}
	return res
}

// @lc code=end
