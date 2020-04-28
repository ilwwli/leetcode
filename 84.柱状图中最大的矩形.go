package main

import "fmt"

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	stack := make([][2]int, 0)
	maxm := 0
	for i := range heights {
		n := len(stack)
		left := i
		for ; n > 0; n-- {
			if stack[n-1][0] >= heights[i] {
				tmp := stack[n-1][0] * (i - stack[n-1][1])
				if tmp > maxm {
					maxm = tmp
				}
				left = stack[n-1][1]
			} else {
				break
			}
		}
		stack = append(stack[:n:n], [2]int{heights[i], left})
	}
	return maxm
}

// @lc code=end
func main() {
	n := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	fmt.Println(n)
}
