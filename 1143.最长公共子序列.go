package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([]int, len(text2)+1)
	for i := 0; i <= len(text2); i++ {
		dp[i] = 0
	}
	for i := 1; i <= len(text1); i++ {
		topleft := dp[0]
		for j := 0; j <= len(text2); j++ {
			if j == 0 {
				dp[j] = 0
			} else if text2[j-1] == text1[i-1] {
				dp[j], topleft = topleft+1, dp[j]
			} else {
				if dp[j-1] < dp[j] {
					dp[j], topleft = dp[j], dp[j]
				} else {
					dp[j], topleft = dp[j-1], dp[j]
				}
			}
		}
	}
	return dp[len(text2)]
}

// @lc code=end
func main() {
	fmt.Println(longestCommonSubsequence("abcde", "ace"))
	fmt.Println(longestCommonSubsequence("abc", "abc"))
	fmt.Println(longestCommonSubsequence("abc", "def"))

}
