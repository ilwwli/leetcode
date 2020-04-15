package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func merging(intervals [][]int) []int {
	start, end := intervals[0][0], intervals[0][1]
	for _, v := range intervals[1:] {
		if v[0] < start {
			start = v[0]
		}
		if v[1] > end {
			end = v[1]
		}
	}
	return []int{start, end}
}

func binarySearchFirstEqualOrGreater(intervals [][]int, target int) int {
	L, R := 0, len(intervals)
	for L < R {
		M := L + (R-L)/2
		if intervals[M][1] < target {
			L = M + 1
		} else {
			R = M
		}
	}
	return L
}

func binarySearchLastEqualOrLess(intervals [][]int, target int) int {
	L, R := 0, len(intervals)
	for L < R {
		M := L + (R-L)/2
		if intervals[M][0] <= target {
			L = M + 1
		} else {
			R = M
		}
	}
	return R - 1
}

func insert(intervals [][]int, newInterval []int) [][]int {
	indleft := binarySearchFirstEqualOrGreater(intervals, newInterval[0])
	indright := binarySearchLastEqualOrLess(
		intervals, newInterval[1])

	left := intervals[:indleft]
	right := intervals[indright+1:]

	needmerge := [][]int{}
	needmerge = append(needmerge, intervals[indleft:indright+1]...)
	needmerge = append(needmerge, newInterval)
	merged := merging(needmerge)

	ans := append(left, append([][]int{merged}, right...)...)
	return ans
}

// @lc code=end

func main() {
	fmt.Println(insert(
		[][]int{
			[]int{1, 5},
		},
		[]int{0, 0},
	))
	fmt.Println(insert(
		[][]int{
			[]int{1, 3},
			[]int{6, 9},
		},
		[]int{2, 5},
	))
	fmt.Println(insert(
		[][]int{
			[]int{1, 2},
			[]int{3, 5},
			[]int{6, 7},
			[]int{8, 10},
			[]int{12, 16},
		},
		[]int{4, 8},
	))
}
