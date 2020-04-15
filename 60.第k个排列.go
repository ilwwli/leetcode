package main

import "fmt"

/*
 * @lc app=leetcode.cn id=60 lang=golang
 *
 * [60] 第k个排列
 */

// @lc code=start
func getPermutation(n int, k int) string {
	factorials := make([]int, n, n)
	factorials[0] = 1
	nums := make([]rune, n, n)
	nums[0] = '1'
	for i := 1; i < n; i++ {
		factorials[i] = factorials[i-1] * i
		nums[i] = nums[i-1] + 1
	}
	ans := make([]rune, n, n)
	k -= 1
	for i := n - 1; i >= 0; i-- {
		div := k / factorials[i]
		k -= factorials[i] * div
		ans[n-1-i] = nums[div]
		nums = append(nums[:div], nums[div+1:]...)
	}
	return string(ans)
}

// @lc code=end
func main() {
	fmt.Println(getPermutation(3, 3))
	fmt.Println(getPermutation(4, 9))
}
