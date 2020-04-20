package main

import "fmt"

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	M, N := len(grid), len(grid[0])
	ans := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == byte('0') {
				continue
			}
			ans++
			queue := [][2]int{[2]int{i, j}}
			for len(queue) > 0 {
				m, n := queue[0][0], queue[0][1]
				queue = queue[1:]
				if grid[m][n] == byte('0') {
					continue
				}
				grid[m][n] = byte('0')
				if m+1 < M && grid[m+1][n] == byte('1') {
					queue = append(queue, [2]int{m + 1, n})
				}
				if m-1 >= 0 && grid[m-1][n] == byte('1') {
					queue = append(queue, [2]int{m - 1, n})
				}
				if n+1 < N && grid[m][n+1] == byte('1') {
					queue = append(queue, [2]int{m, n + 1})
				}
				if n-1 >= 0 && grid[m][n-1] == byte('1') {
					queue = append(queue, [2]int{m, n - 1})
				}
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '1', '0'},
	}
	fmt.Println(numIslands(grid))
}
