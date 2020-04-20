package main

import "fmt"

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	doSubsets(nums, 0, &result, make([]int, 0))
	return result
}

func doSubsets(nums []int, now int, result *[][]int, path []int) {
	if now == len(nums) {
		tmp := append(path[:0:0], path...)
		*result = append(*result, tmp)
		return
	}
	doSubsets(nums, now+1, result, path)
	doSubsets(nums, now+1, result, append(path, nums[now]))
}

// @lc code=end
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
