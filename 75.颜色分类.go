package main

import "fmt"

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	first, last := 0, len(nums)-1
	for now := 0; now <= last; now++ {
		if nums[now] == 0 {
			nums[now], nums[first] = nums[first], nums[now]
			first++
		}
		if nums[now] == 2 {
			nums[now], nums[last] = nums[last], nums[now]
			last--
			now--
		}
	}
	return
}

// @lc code=end
func main() {
	nums := []int{1, 2, 0, 0, 0}
	sortColors(nums)
	fmt.Println(nums)
}
