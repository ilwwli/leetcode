package main

import "fmt"

/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	return doMatch(s, p, make(map[string]bool))
}

func doMatch(s string, p string, cache map[string]bool) bool {
	key := s + "|" + p
	if _, ok := cache[key]; ok {
		return false
	}
	cache[key] = false
	sind := 0
	for pind := 0; pind < len(p); pind++ {
		// deal with *
		if p[pind] == '*' {
			// drop contiguous *
			for pind < len(p) && p[pind] == '*' {
				pind++
			}
			// recursion
			for ; sind <= len(s); sind++ {
				if doMatch(s[sind:], p[pind:], cache) {
					return true
				}
			}
			return false
		}
		// no more string
		if sind >= len(s) {
			return false
		}
		// both advance
		if p[pind] == '?' || p[pind] == s[sind] {
			sind++
			continue
		}
		return false
	}
	return sind == len(s)
}

// @lc code=end
func main() {
	fmt.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))
}
