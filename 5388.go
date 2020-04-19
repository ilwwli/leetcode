package main

import (
	"bytes"
	"fmt"
)

func reformat(s string) string {
	alphas, nums := []rune{}, []rune{}
	for _, v := range s {
		if '0' <= v && v <= '9' {
			nums = append(nums, v)
		} else {
			alphas = append(alphas, v)
		}
	}
	if len(alphas)-len(nums) < -1 || len(alphas)-len(nums) > 1 {
		return ""
	}
	if len(alphas) < len(nums) {
		nums, alphas = alphas, nums
	}
	ans := bytes.NewBufferString("")
	for i := 0; i < len(alphas); i++ {
		ans.WriteRune(alphas[i])
		if i < len(nums) {
			ans.WriteRune(nums[i])
		}
	}
	return ans.String()
}

func main() {
	fmt.Println(reformat("a0b1c2"))
	fmt.Println(reformat("leetcode"))
	fmt.Println(reformat("1229857369"))
	fmt.Println(reformat("covid2019"))
	fmt.Println(reformat("ab123"))
}
