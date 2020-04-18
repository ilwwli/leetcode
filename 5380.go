package main

import (
	"fmt"
	"sort"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	ans := []string{}
	for i := 1; i < len(words); i++ {
		for j := 0; j < i; j++ {
			if isSub(words[j], words[i]) {
				ans = append(ans, words[i])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}
	return ans
}

func isSub(src, tar string) bool {
	i, j := 0, 0
	for i < len(src) && j < len(tar) {
		if src[i] == tar[j] {
			i++
			j++
			if j == len(tar) {
				return true
			}
		} else {
			i = i - j + 1
			j = 0
		}
	}
	return false
}

func main() {
	fmt.Println(stringMatching([]string{}))
}
