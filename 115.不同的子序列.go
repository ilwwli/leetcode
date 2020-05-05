package main

import "fmt"

/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] 不同的子序列
 */

// @lc code=start
func numDistinct(s string, t string) int {
	dp := make([]int, len(s)+1)
	for ind := 1; ind < len(s)+1; ind++ {
		dp[ind] = 1
	}
	last := 1
	for tind := 0; tind < len(t); tind++ {
		for sind := 0; sind < len(s); sind++ {
			if t[tind] == s[sind] {
				dp[sind+1], last = dp[sind]+last, dp[sind+1]
			} else {
				dp[sind+1], last = dp[sind], dp[sind+1]
			}
		}
		//fmt.Println(dp)
		last = dp[0]
	}
	return dp[len(s)]
}

// @lc code=end
func main() {
	//fmt.Println(numDistinct("rabbbit", "rabbit"))
	//fmt.Println(numDistinct("babgbag", "bag"))
	fmt.Println(numDistinct("rddd", "rdd"))
	fmt.Println(numDistinct("ddd", "dd"))
}
