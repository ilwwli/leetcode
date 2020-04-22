package main

import "fmt"

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	doCombine(n, k, 1, &result, make([]int, 0))
	return result
}

func doCombine(n int, k int, now int, result *[][]int, path []int) {
	if n-now+1 < k {
		return
	}
	if k == 0 {
		tmp := append(path[:0:0], path...)
		*result = append(*result, tmp)
		return
	}
	path = append(path, -1)
	for i := now; i <= n; i++ {
		path[len(path)-1] = i
		doCombine(n, k-1, i+1, result, path)
	}
	path = path[:len(path)-1]
	return
}

// @lc code=end
func main() {
	fmt.Println(combine(4, 2))
}
