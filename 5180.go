package main

import (
	"container/list"
	"fmt"
)

func constrainedSubsetSum(nums []int, k int) int {
	start := 0
	maxm := nums[0]
	for ; start < len(nums); start++ {
		if nums[start] >= 0 {
			break
		}
		if nums[start] > maxm {
			maxm = nums[start]
		}
	}
	if start == len(nums) {
		return maxm
	}
	nums = nums[start:]

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxm = nums[0]
	dq := list.New()
	for i := 1; i < len(nums); i++ {
		// pushback i-1
		for node := dq.Back(); node != nil; {
			if node.Value.(int) >= dp[i-1] {
				break
			}
			node = node.Prev()
			dq.Remove(dq.Back())
		}
		dq.PushBack(dp[i-1])
		// pop i-k
		if i-1 >= k && dq.Front().Value.(int) == dp[i-1-k] {
			dq.Remove(dq.Front())
		}
		// getmax
		tmp := 0
		if dq.Front() != nil && dq.Front().Value.(int) > 0 {
			tmp = dq.Front().Value.(int)
		}
		dp[i] = nums[i] + tmp
		if dp[i] > maxm {
			maxm = dp[i]
		}
	}
	return maxm
}

func main() {
	fmt.Println(constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2))
	fmt.Println(constrainedSubsetSum([]int{-1, -2, -3}, 1))
	fmt.Println(constrainedSubsetSum([]int{10, -2, -10, -5, 20}, 2))
	fmt.Println(constrainedSubsetSum([]int{-7609, 249, -1699, 2385, 9125, -9037, 1107, -8228, 856, -9526}, 9))
}
