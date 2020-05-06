package main

import "fmt"

/*
 * @lc app=leetcode.cn id=983 lang=golang
 *
 * [983] 最低票价
 */

// @lc code=start
func mincostTickets(days []int, costs []int) int {
	dp := make([]int, 366)
	dayind := 0
	for i := 1; i <= 365; i++ {
		if dayind >= len(days) {
			break
		}
		if i == days[dayind] {
			dp[i] = min(
				dp[i-1]+costs[0],
				dp[max(0, i-7)]+costs[1],
				dp[max(0, i-30)]+costs[2],
			)
			dayind++
		} else {
			dp[i] = dp[i-1]
		}

	}
	//fmt.Println(dp)
	return dp[days[dayind-1]]
}

func min(i ...int) int {
	minm := i[0]
	for _, v := range i {
		if v < minm {
			minm = v
		}
	}
	return minm
}

func max(i ...int) int {
	maxm := i[0]
	for _, v := range i {
		if v > maxm {
			maxm = v
		}
	}
	return maxm
}

// @lc code=end
func main() {
	fmt.Println(mincostTickets([]int{
		1, 4, 6, 7, 8, 20,
	}, []int{
		2, 7, 15,
	}))
	fmt.Println(mincostTickets([]int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31,
	}, []int{
		2, 7, 15,
	}))
}
