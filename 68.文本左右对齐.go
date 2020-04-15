package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=68 lang=golang
 *
 * [68] 文本左右对齐
 */

// @lc code=start
func fullJustify(words []string, maxWidth int) []string {
	ans := []string{}
	tmp := []string{}
	nowWidth := -1
	for _, word := range words {
		if len(word)+1+nowWidth <= maxWidth {
			tmp = append(tmp, word)
			nowWidth += 1 + len(word)
		} else {
			ans = append(ans, insertSpaces(tmp, maxWidth))
			tmp = []string{word}
			nowWidth = len(word)
		}
	}
	ans = append(ans, insertSpaces([]string{strings.Join(tmp, " ")}, maxWidth))
	return ans
}

func insertSpaces(words []string, maxWidth int) string {
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", maxWidth-len(words[0]))
	}
	needspaces := maxWidth - len(words) + 1
	for _, word := range words {
		needspaces -= len(word)
	}
	quotient, remainder := needspaces/(len(words)-1), needspaces%(len(words)-1)
	for i := 0; i < remainder; i++ {
		words[i] += " "
	}
	return strings.Join(words, strings.Repeat(" ", quotient+1))
}

// @lc code=end
func main() {
	ans := fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	for _, a := range ans {
		fmt.Println(a)
	}
}
