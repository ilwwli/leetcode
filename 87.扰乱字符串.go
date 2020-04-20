package main

import "fmt"

/*
 * @lc app=leetcode.cn id=87 lang=golang
 *
 * [87] 扰乱字符串
 */

// @lc code=start
func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return doIsScramble(s1, s2, make(map[string]bool))
}

func doIsScramble(s1, s2 string, cache map[string]bool) bool {
	if ans, ok := cache[s1+"|"+s2]; ok {
		return ans
	}
	rev := reverse(s1)
	if s1 == s2 || rev == s2 {
		cache[s1+"|"+s2] = true
		return true
	}
	for i := 1; i < len(s1); i++ {
		if doIsScramble(s1[i:], s2[i:], cache) && doIsScramble(s1[:i], s2[:i], cache) {
			cache[s1+"|"+s2] = true
			return true
		}
		if doIsScramble(rev[i:], s2[i:], cache) && doIsScramble(rev[:i], s2[:i], cache) {
			cache[s1+"|"+s2] = true
			return true
		}
	}
	cache[s1+"|"+s2] = false
	return false
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// @lc code=end
func main() {
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abcde", "caebd"))
	fmt.Println(isScramble("abc", "bac"))
	fmt.Println(isScramble("abc", "bca"))
}
