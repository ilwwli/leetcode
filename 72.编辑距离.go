package main

import "fmt"

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	dp := make([]int, len(word2))
	for i := 0; i < len(word2); i++ {
		dp[i] = i + 1
	}
	var topleft, left, top int
	for row, v1 := range word1 {
		topleft, left = row, row+1
		for column, v2 := range word2 {
			top = dp[column]
			if v1 == v2 {
				topleft--
			}
			dp[column] = 1 + min(left, top, topleft)
			topleft, left = top, dp[column]
		}
	}
	return dp[len(word2)-1]
}

func min(nums ...int) int {
	m := nums[0]
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}

// @lc code=end
func main() {
	fmt.Println(minDistance("intention", "execution"))
}
