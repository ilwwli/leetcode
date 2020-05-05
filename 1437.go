package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	last := -k - 1
	for ind, v := range nums {
		if v == 0 {
			continue
		}
		if ind-last < k+1 {
			return false
		}
		last = ind
	}
	return true
}

func main() {
	fmt.Println(kLengthApart(
		[]int{1, 0, 0, 1, 0, 1}, 2,
	))
	fmt.Println(kLengthApart(
		[]int{1, 1, 1, 1, 1}, 0,
	))
	fmt.Println(kLengthApart(
		[]int{}, 1,
	))
}
