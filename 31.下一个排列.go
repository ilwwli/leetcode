package main

import "fmt"

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
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

// @lc code=end

func main() {
	l := []int{1, 1, 5, 5, 5}
	nextPermutation(l)
	fmt.Println(l)
}
