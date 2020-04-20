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

