package main

import (
	"fmt"
)

func minTime(time []int, m int) int {
	left, right := 0, 1000000000 // 10^5 * 10^4 = 10^9
	for left < right {
		mid := left + (right-left)/2
		if !judge(time, m, mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func judge(time []int, m int, T int) bool {
	summ := 0
	maxm := 0
	for _, v := range time {
		summ += v
		if v > maxm {
			maxm = v
		}
		if summ-maxm > T {
			m--
			if m <= 0 {
				return false
			}
			summ = v
			maxm = v
		}
	}
	return true
}

func main() {
	fmt.Println(minTime([]int{1, 2, 3, 3}, 2))
	fmt.Println(minTime([]int{1, 2, 3, 3, 3}, 2))
	fmt.Println(minTime([]int{50, 47, 68, 33, 35, 84, 25, 49, 91, 75}, 1))
}
