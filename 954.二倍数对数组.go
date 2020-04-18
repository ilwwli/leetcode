package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=954 lang=golang
 *
 * [954] 二倍数对数组
 */

// @lc code=start
func canReorderDoubled(A []int) bool {
	sort.Sort(sort.IntSlice(A))
	m := make(map[int]int)
	for _, v := range A {
		var key int
		if v >= 0 {
			if v%2 == 1 {
				m[v]++
				continue
			}
			key = v / 2
		} else {
			key = v * 2
		}
		if m[key] > 0 {
			m[key]--
		} else {
			m[v]++
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
func main() {
	fmt.Println(canReorderDoubled([]int{4, -2, 2, -4}))
	fmt.Println(canReorderDoubled([]int{2, 2}))
}
