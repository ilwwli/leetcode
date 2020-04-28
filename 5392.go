package main

import "fmt"

func maxScore(s string) int {
	left := 0
	if s[0] == '0' {
		left++
	}
	if s[len(s)-1] == '1' {
		left++
	}
	s = s[1 : len(s)-1]
	right := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			right++
		}
	}
	maxm := right
	now := right
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			now++
			if now > maxm {
				maxm = now
			}
		} else {
			now--
		}
	}
	return maxm + left
}

func main() {
	fmt.Println(maxScore("011101"))
	fmt.Println(maxScore("00111"))
	fmt.Println(maxScore("1111"))
	fmt.Println(maxScore("01"))
	fmt.Println(maxScore("10"))
	fmt.Println(maxScore("0000"))
}
