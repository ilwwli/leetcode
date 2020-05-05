package main

import (
	"container/list"
	"fmt"
)

func longestSubarray(nums []int, limit int) int {
	lastpop, maxlength := -1, 0
	minlist, maxlist := list.New(), list.New()
	for ind, v := range nums {
		// push to stack
		for minlist.Len() > 0 {
			bind := minlist.Back().Value.(int)
			if v >= nums[bind] {
				break
			}
			minlist.Remove(minlist.Back())
		}
		minlist.PushBack(ind)
		// deal with stack
		for minlist.Len() > 0 {
			find := minlist.Front().Value.(int)
			if v-nums[find] <= limit {
				break
			}
			minlist.Remove(minlist.Front())
			if find > lastpop {
				lastpop = find
			}
		}

		// push to stack
		for maxlist.Len() > 0 {
			bind := maxlist.Back().Value.(int)
			if v <= nums[bind] {
				break
			}
			maxlist.Remove(maxlist.Back())
		}
		maxlist.PushBack(ind)
		// deal with stack
		for maxlist.Len() > 0 {
			find := maxlist.Front().Value.(int)
			if nums[find]-v <= limit {
				break
			}
			maxlist.Remove(maxlist.Front())
			if find > lastpop {
				lastpop = find
			}
		}

		if ind-lastpop > maxlength {
			maxlength = ind - lastpop
		}
	}
	return maxlength
}

func main() {
	fmt.Println(longestSubarray(
		[]int{8, 2, 4, 7},
		4,
	))
	fmt.Println(longestSubarray(
		[]int{10, 1, 2, 4, 7, 2},
		5,
	))
	fmt.Println(longestSubarray(
		[]int{4, 2, 2, 2, 4, 4, 2, 2},
		0,
	))
}
