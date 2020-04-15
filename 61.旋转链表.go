/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	len := 0
	fast, slow := head, head
	for i:=0;i<k;i++ {
		fast = fast.Next
		len++
		if fast == nil {
			k = k % len + len
			fast = head
		}
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fast.Next = head
	tmp := slow.Next
	slow.Next = nil
	return tmp
}
// @lc code=end

