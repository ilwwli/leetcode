package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	if target <= 0 {
		return nil
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] > candidates[j]
	})
	return doCombSum(candidates, target, 0, make(map[[2]int][][]int))
}

func doCombSum(candidates []int, target int, index int, cache map[[2]int][][]int) [][]int {
	// boundary 1
	if target == 0 {
		return [][]int{[]int{}}
	}
	// skip useless candidates in this recursion
	for index < len(candidates) && candidates[index] > target {
		index++
	}
	// boundary 2
	if index >= len(candidates) {
		return [][]int{}
	}
	// try cache
	if cache[[2]int{target, index}] != nil {
		return cache[[2]int{target, index}]
	}
	// do
	ans := [][]int{}
	next := index + 1
	for next < len(candidates) && candidates[next] == candidates[index] {
		next++
	}
	now := []int{}
	for d := 0; d <= next-index && target-d*candidates[index] >= 0; d++ {
		for _, v := range doCombSum(candidates, target-d*candidates[index], next, cache) {
			ans = append(ans, append(v, now...))
		}
		now = append(now, candidates[index])
	}
	cache[[2]int{target, index}] = ans
	return ans
}

// @lc code=end

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
	fmt.Println(combinationSum2([]int{}, 5))
	fmt.Println(combinationSum2([]int{}, 0))
}
