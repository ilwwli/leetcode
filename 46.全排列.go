package main

import "fmt"

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	ans = append(ans, append([]int{}, nums...))
	times := 1
	for i := 1; i <= len(nums); i++ {
		times *= i
	}
	for i := 1; i < times; i++ {
		nextPermutation(nums)
		ans = append(ans, append([]int{}, nums...))
	}
	return ans
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

// @lc code=end
func main() {
	fmt.Println(permute([]int{3, 2, 1}))
}
