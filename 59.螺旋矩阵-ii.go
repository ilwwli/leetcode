package main

import "fmt"

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n, n)
	}
	top, bottom, left, right := -1, n, -1, n
	state := 1
	i, j := 0, 0
	index := 1
	for {
		if bottom-top <= 1 || right-left <= 1 {
			break
		}
		switch state {
		case 1:
			for j = left + 1; j < right; j++ {
				matrix[i][j] = index
				index++
			}
			j--
			state = 2
			top++

		case 2:
			for i = top + 1; i < bottom; i++ {
				matrix[i][j] = index
				index++
			}
			i--
			state = 3
			right--
		case 3:
			for j = right - 1; j > left; j-- {
				matrix[i][j] = index
				index++
			}
			j++
			state = 4
			bottom--
		case 4:
			for i = bottom - 1; i > top; i-- {
				matrix[i][j] = index
				index++
			}
			i++
			state = 1
			left++
		}
	}
	return matrix
}

// @lc code=end
func main() {
	for n := 1; n < 5; n++ {
		fmt.Println(generateMatrix(n))
	}
}
