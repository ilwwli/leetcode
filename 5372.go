package main

import "fmt"

func minStartValue(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	minm := nums[0]
	now := 0
	for _, v := range nums {
		now += v
		minm = min(minm, now)
	}
	if minm >= 0 {
		return 1
	} else {
		return -minm + 1
	}
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	fmt.Println(minStartValue([]int{-3, 2, -3, 4, 2}))
	fmt.Println(minStartValue([]int{1, -2, -3}))
}
