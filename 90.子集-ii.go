package main

import "fmt"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	counter := map[int]int{}
	for _, v := range nums {
		counter[v]++
	}
	nums = nil
	for k := range counter {
		nums = append(nums, k)
	}
	ans := make([][]int, 0)
	var f func(int, []int)
	f = func(ind int, tmp []int) {
		if ind >= len(nums) {
			ans = append(ans, tmp)
			return
		}
		for i := 0; i <= counter[nums[ind]]; i++ {
			f(ind+1, append(tmp[:0:0], tmp...))
			tmp = append(tmp, nums[ind])
		}
	}
	f(0, make([]int, 0))
	return ans
}

// @lc code=end
func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
