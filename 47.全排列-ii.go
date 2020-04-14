package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {

	ans := make([][]int, 0)
	sort.Sort(sort.IntSlice(nums))
	ans = append(ans, append([]int{}, nums...))
	for nextPermutation(nums) {
		ans = append(ans, append([]int{}, nums...))
	}
	return ans
}

func nextPermutation(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == 0 {
			return false
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
	return true
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
	fmt.Println(permuteUnique([]int{2, 1, 1}))
}
