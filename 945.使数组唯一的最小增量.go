package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=945 lang=golang
 *
 * [945] 使数组唯一的最小增量
 */

// @lc code=start
func minIncrementForUnique(A []int) int {
	sort.Ints(A)
	taken := 0
	steps := 0
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			taken++
			steps -= A[i]
		} else if drops := A[i] - A[i-1] - 1; drops > 0 {
			if drops > taken {
				drops = taken
			}
			taken -= drops
			steps += (A[i-1]*2 + 1 + drops) * drops / 2
		}
	}
	if taken > 0 {
		steps += (A[len(A)-1]*2 + 1 + taken) * taken / 2
	}
	return steps
}

// @lc code=end

func main() {
	fmt.Println(minIncrementForUnique([]int{1, 2, 1, 1, 1, 10}))
}
