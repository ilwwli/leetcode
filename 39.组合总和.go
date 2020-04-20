package main

import "fmt"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	return dfs(candidates, target, 0, )
}

func dfs(candidates []int, remain int, pos int, paths) [][]int {
	// 边界条件:找到了，返回含有一行的二维数组
	if remain == 0 {
		return [][]int{[]int{}}
	}
	// 边界条件:错误情况，返回空二维数组
	if pos >= len(candidates) || remain < 0 {
		return [][]int{}
	}
	// 减少重复子问题
	if cache[[2]int{remain, pos}] != nil {
		return cache[[2]int{remain, pos}]
	}
	// 返回值
	ans := [][]int{}
	// 缓存
	tmp := []int{}
	for remain >= 0 {
		for _, v := range dfs(candidates, remain, pos+1, cache) {
			ans = append(ans, append(v, tmp...))
		}
		tmp = append(tmp, candidates[pos])
		remain -= candidates[pos]
	}
	cache[[2]int{remain, pos}] = ans
	return ans
}

// @lc code=end
func main() {
	fmt.Println(combinationSum(
		[]int{2, 3, 5}, 8))
}
