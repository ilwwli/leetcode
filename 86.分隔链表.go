package main

import "fmt"

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	sentinalHead := &ListNode{
		Next: head,
	}
	sendtinalBig := &ListNode{
		Next: nil,
	}
	pre, curr := sentinalHead, head
	big := sendtinalBig
	for ; curr != nil; curr = curr.Next {
		if curr.Val >= x {
			big.Next = curr
			big = big.Next
			pre.Next = curr.Next
		} else {
			pre = pre.Next
		}
	}
	// Must Added to avoid circular list
	big.Next = nil
	pre.Next = sendtinalBig.Next
	return sentinalHead.Next
}

// @lc code=end
func main() {
	root := sliceToListNode([]int{1, 4, 3, 2, 5, 2})
	fmt.Println(listNodeToSlice(root))
	root = partition(root, 3)
	fmt.Println(listNodeToSlice(root))
}

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
