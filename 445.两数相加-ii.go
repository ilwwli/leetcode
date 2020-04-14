package main

import (
	"math/big"
)

// 还不如用栈做。虽然节省了一点空间，bigInt效率太低了。

/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b := big.NewInt(0), big.NewInt(0)
	for l1 != nil {
		a.Mul(a, big.NewInt(10))
		a.Add(a, big.NewInt(int64(l1.Val)))
		l1 = l1.Next
	}
	for l2 != nil {
		b.Mul(b, big.NewInt(10))
		b.Add(b, big.NewInt(int64(l2.Val)))
		l2 = l2.Next
	}
	a = a.Add(a, b)
	phony := &ListNode{Next: &ListNode{}}
	tmp := phony
	for _, v := range a.String() {
		tmp = tmp.Next
		tmp.Val = int(v - '0')
		tmp.Next = &ListNode{}
	}
	tmp.Next = nil
	return phony.Next
}

// @lc code=end
