package main

import "fmt"

/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	if n <= 0 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 0; i <= n; i++ {
		for j := 0; j <= i-1; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

// @lc code=end
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(numTrees(i))
	}
}
