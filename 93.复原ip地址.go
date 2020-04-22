package main

import "fmt"

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原IP地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	for i := 1; i <= 3; i++ {
		for j := i + 1; j <= i+3; j++ {
			for k := j + 1; k <= j+3; k++ {
				if k < len(s) && isValid(s[0:i]) && isValid(s[i:j]) && isValid(s[j:k]) && isValid(s[k:]) {
					ans = append(ans, s[0:i]+"."+s[i:j]+"."+s[j:k]+"."+s[k:])
				}
			}
		}
	}
	return ans
}

func isValid(s string) bool {
	if len(s) < 1 || len(s) > 3 {
		return false
	}
	if len(s) == 1 {
		return true
	} else if len(s) == 2 {
		return s[0] != byte('0')
	} else {
		if s[0] == byte('1') {
			return true
		} else if s[0] == byte('2') {
			if s[1] <= byte('4') {
				return true
			} else if s[1] == byte('5') {
				return s[2] <= byte('5')
			} else {
				return false
			}
		} else {
			return false
		}
	}
}

// @lc code=end
func main() {
	fmt.Println(restoreIpAddresses("2552551113"))
}
