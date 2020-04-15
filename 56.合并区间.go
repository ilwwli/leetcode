package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
type Endpoint struct {
	Val  int
	Type bool // true->start, false->end
}

type EndpointSlice []Endpoint

func (s EndpointSlice) Len() int { return len(s) }
func (s EndpointSlice) Less(i, j int) bool {
	if s[i].Val < s[j].Val || (s[i].Val == s[j].Val && s[i].Type && !s[j].Type) {
		return true
	}
	return false
}
func (s EndpointSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	s := EndpointSlice{}
	for _, interval := range intervals {
		s = append(s, Endpoint{interval[0], true}, Endpoint{interval[1], false})
	}
	sort.Sort(s)
	state := 0
	start := -1
	for _, ep := range s {
		if ep.Type {
			state++
		} else {
			state--
		}
		if state == 0 {
			ans = append(ans, []int{start, ep.Val})
			start = -1
		} else {
			if start == -1 {
				start = ep.Val
			}
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(merge([][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}))
	fmt.Println(merge([][]int{
		[]int{1, 4},
		[]int{4, 5},
	}))
	fmt.Println(merge([][]int{}))
}
