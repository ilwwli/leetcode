package main

import "fmt"

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	N := len(matrix)
	if N == 0 {
		return
	}
	for i := 0; i < (N+1)/2; i++ {
		for j := 0; j < N/2; j++ {
			matrix[i][j], matrix[j][N-1-i], matrix[N-1-i][N-1-j], matrix[N-1-j][i] = matrix[N-1-j][i], matrix[i][j], matrix[j][N-1-i], matrix[N-1-i][N-1-j]
		}
	}
}

// @lc code=end
func main() {
	matrix := [][]int{
		[]int{5, 1, 9, 11},
		[]int{2, 4, 8, 10},
		[]int{13, 3, 6, 7},
		[]int{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
