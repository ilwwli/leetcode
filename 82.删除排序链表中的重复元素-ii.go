/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	sentinal := &ListNode{
		Next: head,
	}
	sentinal2 := &ListNode{
		Next: sentinal,
	}
	slow, fast := sentinal2, sentinal.Next
	tmp := head.Val - 1
	flag := false
	for ;fast != nil;fast = fast.Next {		
		if fast.Val == tmp {
			flag = true
			continue
		}
		if !flag {			
			slow = slow.Next
		} else {
			slow.Next = fast
		}
		tmp = fast.Val
		flag = false		
	}
	if flag {
		slow.Next = nil
	}
	return sentinal.Next
}

// @lc code=end
