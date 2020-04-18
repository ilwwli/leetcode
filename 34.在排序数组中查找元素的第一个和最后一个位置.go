package main

import "fmt"

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] >= target {
			end = mid
		}
	}
	if nums[start] == target {
		res[0] = start
	} else {
		return res
	}
	start, end = 0, len(nums)-1
	for start < end {
		mid := start + (end-start+1)/2
		if nums[mid] <= target {
			start = mid
		} else if nums[mid] > target {
			end = mid - 1
		}
	}
	if nums[start] == target {
		res[1] = start
	}
	return res
}

// @lc code=end

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
}
