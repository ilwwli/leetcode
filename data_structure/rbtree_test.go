package data_structure

import (
	"testing"
)

func TestInsertAndDelete(t *testing.T) {
	tree := &RBTree{}
	values := make([]int, 100)
	for i := 1; i <= 100; i++ {
		values[i-1] = i
	}
	for i := 0; i < 100; i++ {
		for _, v := range values {
			tree.Insert(v)
			propJudge(tree, t)
		}
		for _, v := range values {
			tree.Delete(v)
			propJudge(tree, t)
		}
		nextPermutation(values)
	}
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == 0 {
			reverse(nums, 0, len(nums)-1)
			break
		}
		if nums[i-1] >= nums[i] {
			continue
		}
		for j := len(nums) - 1; j > i-1; j-- {
			if nums[j] > nums[i-1] {
				nums[j], nums[i-1] = nums[i-1], nums[j]
				break
			}
		}
		reverse(nums, i, len(nums)-1)
		break
	}
	return
}

func reverse(nums []int, start, end int) {
	if end >= len(nums) || start < 0 {
		return
	}
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func propJudge(tree *RBTree, t *testing.T) bool {
	if tree.Root != nil && !tree.Root.Black {
		PrintTree(tree)
		t.Fatal("root is not black")
		return false
	}
	m := make(map[int]bool)
	if !judgeHelper(tree.Root, true, m, 0) {
		PrintTree(tree)
		t.Fatal("color not right")
		return false
	}
	if len(m) != 1 {
		PrintTree(tree)
		t.Fatal("black height not same")
		return false
	}
	return true
}

func judgeHelper(node *RBTreeNode, blackParent bool, height map[int]bool, tmpheight int) bool {
	if node == nil {
		height[tmpheight] = true
		return true
	}
	if node.Black == false && blackParent == false {
		return false
	}
	if node.Black {
		tmpheight++
	}
	return judgeHelper(node.Left, node.Black, height, tmpheight) && judgeHelper(node.Right, node.Black, height, tmpheight)
}
