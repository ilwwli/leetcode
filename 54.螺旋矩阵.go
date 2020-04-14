package main

import "fmt"

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	top, bottom, left, right := -1, len(matrix), -1, len(matrix[0])
	ans := []int{}
	state := 1
	i, j := 0, 0
	for {
		if bottom-top <= 1 || right-left <= 1 {
			break
		}
		switch state {
		case 1:
			for j = left + 1; j < right; j++ {
				ans = append(ans, matrix[i][j])
			}
			j--
			state = 2
			top++
		case 2:
			for i = top + 1; i < bottom; i++ {
				ans = append(ans, matrix[i][j])
			}
			i--
			state = 3
			right--
		case 3:
			for j = right - 1; j > left; j-- {
				ans = append(ans, matrix[i][j])
			}
			j++
			state = 4
			bottom--
		case 4:
			for i = bottom - 1; i > top; i-- {
				ans = append(ans, matrix[i][j])
			}
			i++
			state = 1
			left++
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(spiralOrder([][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{0, 10, 11, 12},
	}))
}
