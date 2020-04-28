/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) <= 0 {
        return nil
    }
    for n:=1;n<len(lists);n *= 2 {
        for i:=0;i <len(lists);i += n * 2 {
            j := i + n
            if j >= len(lists) {
                break
            }           
            lists[i] = mergeTwoLists(lists[i], lists[j])
        }
    }
    return lists[0]
}

func mergeTwoLists(root1, root2 *ListNode) *ListNode{
    sentinal := &ListNode{}
    ptr1, ptr2, now := root1, root2, sentinal
    for ptr1 != nil && ptr2 != nil {        
        if ptr1.Val < ptr2.Val {
            now.Next = ptr1
            now = now.Next
            ptr1 = ptr1.Next
        } else {
            now.Next = ptr2
            now = now.Next
            ptr2 = ptr2.Next
        }
    }
    if ptr1 != nil {
        now.Next = ptr1
    } else {
        now.Next = ptr2
    }
    return sentinal.Next
}
// @lc code=end

