package main

import "fmt"

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}
	ans, start, end := 0, 0, len(height)-1
	maxstart, maxend := height[start], height[end]
	for start != end {
		for maxstart <= maxend && start != end {
			start++
			if height[start] <= maxstart {
				ans += maxstart - height[start]
			} else {
				maxstart = height[start]
			}
		}
		for maxstart > maxend && start != end {
			end--
			if height[end] <= maxend {
				ans += maxend - height[end]
			} else {
				maxend = height[end]
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(trap([]int{5, 5, 5, 5, 5}))
}
