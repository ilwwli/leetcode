package main

import "fmt"

/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 */

// @lc code=start
func isNumber(s string) bool {
	jumptable := [][]int{
		[]int{0, 1, 6, 2, -1, -1},
		[]int{-1, -1, 6, 2, -1, -1},
		[]int{-1, -1, 3, -1, -1, -1},
		[]int{8, -1, 3, -1, 4, -1},
		[]int{-1, 7, 5, -1, -1, -1},
		[]int{8, -1, 5, -1, -1, -1},
		[]int{8, -1, 6, 3, 4, -1},
		[]int{-1, -1, 5, -1, -1, -1},
		[]int{8, -1, -1, -1, -1, -1},
	}
	currstate := 0
	for _, v := range s {
		if currstate == -1 {
			return false
		}
		var index int
		switch {
		case v == ' ':
			index = 0
		case v == '+':
			fallthrough
		case v == '-':
			index = 1
		case v >= '0' && v <= '9':
			index = 2
		case v == '.':
			index = 3
		case v == 'e':
			index = 4
		default:
			index = 5
		}
		currstate = jumptable[currstate][index]
	}
	if currstate == 3 || currstate == 5 || currstate == 6 || currstate == 8 {
		return true
	}
	return false
}

// @lc code=end
func main() {
	fmt.Println(isNumber(" 0.1 "))
}
