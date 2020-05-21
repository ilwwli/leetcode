package main

import "fmt"

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, dummy
	for fast != nil {
		slow = fast
		for i := 0; i < k && fast != nil; i++ {
			fast = fast.Next
		}
		if fast == nil {
			return dummy.Next
		}
		fast = reverseList(slow, fast)
	}
	return dummy.Next
}

func reverseList(start, end *ListNode) *ListNode {
	slow, fast := start, start.Next
	before, last := start.Next, end.Next
	for slow != end {
		slow, fast, fast.Next = fast, fast.Next, slow
	}
	start.Next, start.Next.Next = end, last
	return before
}

// @lc code=end
type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToListNode(s []int) *ListNode {
	sentinal := &ListNode{
		Next: nil,
	}
	curr := sentinal
	for _, v := range s {
		tmp := &ListNode{
			Val:  v,
			Next: nil,
		}
		curr.Next = tmp
		curr = curr.Next
	}
	return sentinal.Next
}

func listNodeToSlice(root *ListNode) []int {
	ans := make([]int, 0)
	for ; root != nil; root = root.Next {
		ans = append(ans, root.Val)
	}
	return ans
}
func main() {
	ls := sliceToListNode([]int{1, 2, 3, 4, 5})
	ans := reverseKGroup(ls, 2)
	fmt.Println(listNodeToSlice(ans))
}
