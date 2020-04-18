package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	maxl := 0
	left, right := 0, 0
	for _, v := range s {
		if v == '(' {
			left++
			continue
		}
		right++
		if right == left {
			maxl = max(maxl, left+right)
		} else if right > left {
			left, right = 0, 0
		}

	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		v := s[i]
		if v == ')' {
			left++
			continue
		}
		right++
		if right == left {
			maxl = max(maxl, left+right)
		} else if right > left {
			left, right = 0, 0
		}

	}
	return maxl
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

func main() {
	fmt.Println(longestValidParentheses(")(((((()())()()))()(()))("))
	fmt.Println(longestValidParentheses(")()()"))
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(""))
}
