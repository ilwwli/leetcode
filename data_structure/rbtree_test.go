package data_structure

import (
	"fmt"
	"testing"
)

func TestInsertAndDelete(t *testing.T) {
	tree := &RBTree{}
	for i := 1; i < 10; i++ {
		tree.Insert(i)
	}
	PrintTree(tree)
	for i := 1; i < 10; i++ {
		tree.Delete(i)
	}
	PrintTree(tree)
}

func PrintTree(tree *RBTree) {
	list := make([]*RBTreeNode, 0)
	list = append(list, tree.Root)
	ans := make([]int, 0)
	ans2 := make([]bool, 0)
	for len(list) > 0 {
		top := list[0]
		list = list[1:]
		if top != nil {
			ans = append(ans, top.Val)
			ans2 = append(ans2, top.Black)
			list = append(list, top.Left, top.Right)
		} else {
			ans = append(ans, -1)
			ans2 = append(ans2, true)
		}
	}
	fmt.Println(ans)
	fmt.Println(ans2)
}

// 			4
// 	2r 				6r
// 1 		3 		5 		8
// -1 -1   -1 -1   -1 -1  7r  9r
