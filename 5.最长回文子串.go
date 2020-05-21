package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	// preprocess
	prebuf := strings.Builder{}
	prebuf.WriteRune('#')
	for _, v := range s {
		prebuf.WriteRune(v)
		prebuf.WriteRune('#')
	}
	s = prebuf.String()

	max_arm_index := 0
	arm_len := make([]int, len(s))
	right, center := -1, -1
	for i := 1; i < len(s); i++ {
		var cur_arm_len int
		if right < i {
			cur_arm_len = expand(s, i, i)
		} else {
			i_sym := 2*center - i
			min_arm_len := min(arm_len[i_sym], right-i)
			cur_arm_len = expand(s, i-min_arm_len, i+min_arm_len)
		}
		arm_len[i] = cur_arm_len
		if i+cur_arm_len > right {
			right = i + cur_arm_len
			center = i
		}
		if cur_arm_len > arm_len[max_arm_index] {
			max_arm_index = i
		}
	}
	//fmt.Println(arm_len)
	ans := ""
	left, right := max_arm_index-arm_len[max_arm_index], max_arm_index+arm_len[max_arm_index]
	for _, v := range s[left : right+1] {
		if v != '#' {
			ans += string(v)
		}
	}
	return ans
}

func expand(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return (right - left - 2) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
func main() {
	fmt.Println(longestPalindrome("babad"))	
}
