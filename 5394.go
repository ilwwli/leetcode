package main

import (
	"container/list"
	"fmt"
)

func findDiagonalOrder(nums [][]int) []int {
	ans := make([]int, 0, len(nums))
	ls := list.New()
	ls.PushBack([2]int{0, 0})
	if len(nums) > 1 {
		ls.PushBack([2]int{-1, 1})
	}
	for ls.Len() > 0 {
		ele := ls.Front().Value.([2]int)
		i, j := ele[0], ele[1]
		ls.Remove(ls.Front())
		if i == -1 {
			ls.PushFront([2]int{j, 0})
			if j < len(nums)-1 {
				ls.PushBack([2]int{i, j + 1})
			}
			continue
		}
		ans = append(ans, nums[i][j])
		if j < len(nums[i])-1 {
			ls.PushBack([2]int{i, j + 1})
		}
	}
	return ans
}

func main() {
	square := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7},
		{8},
		{9, 10, 11},
		{12, 13, 14, 15, 16},
	}
	fmt.Println(findDiagonalOrder(square))
	square = [][]int{
		{1, 2, 3, 4, 5, 6},
	}
	fmt.Println(findDiagonalOrder(square))
	square = [][]int{
		{1},
		{2},
		{3},
	}
	fmt.Println(findDiagonalOrder(square))
}
