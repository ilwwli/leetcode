package main

import "fmt"

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	first, last, now := 1, 1, 0
	for i := 0; i < len(s); i++ {
		if i >= 1 && (s[i-1] == byte('1') || s[i-1] == byte('2') && s[i] <= byte('6')) {
			now += first
		}
		if s[i] != byte('0') {
			now += last
		}
		first, last, now = last, now, 0
	}
	return last
}

// @lc code=end
func main() {
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("00"))
	fmt.Println(numDecodings("226"))
}
