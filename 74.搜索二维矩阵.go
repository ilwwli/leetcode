package main

import "fmt"

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	front := make([]int, 0, m)
	for _, v := range matrix {
		front = append(front, v[0])
	}
	left := binarySearchLessOrEqual(front, target)
	if left < 0 || matrix[left][n-1] < target {
		return false
	}
	mid := binarySearchLessOrEqual(matrix[left], target)
	if matrix[left][mid] != target {
		return false
	}
	return true
}

func binarySearchLessOrEqual(l []int, target int) int {
	if len(l) <= 0 {
		return -1
	}
	start, end := 0, len(l)-1
	for start < end {
		mid := start + (end-start)/2
		if l[mid] <= target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if l[start] > target {
		start--
	}
	return start
}

// @lc code=end
func main() {
	fmt.Println(binarySearchLessOrEqual([]int{1}, 0))
}
