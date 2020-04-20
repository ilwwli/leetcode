package main

import "fmt"

/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	tmp := nums[0] - 1
	dup := false
	for _, v := range nums {
		if v == tmp && dup {
			continue
		}
		if v == tmp {
			dup = true
		} else {
			tmp = v
			dup = false
		}
		nums[slow] = v
		slow++
	}
	return slow
}

// @lc code=end
func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	ind := removeDuplicates(nums)
	fmt.Println(ind, nums[:ind])
}
