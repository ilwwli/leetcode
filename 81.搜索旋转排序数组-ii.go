package main

import "fmt"

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start
func search(nums []int, target int) bool {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] {
			left++
			continue
		}

		if nums[left] < nums[mid] { // front part ordered
			if nums[left] <= target && target < nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else { // back part ordered
			if nums[mid] < target && target <= nums[right-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return false
}

// @lc code=end
func main() {
	fmt.Println(search([]int{
		7, 6, 7, 7, 7, 7, 7,
	}, 6))
	fmt.Println(search([]int{
		7, 7, 7, 7, 6, 7,
	}, 6))
	fmt.Println(search([]int{
		7, 7, 7, 7, 7, 7, 7,
	}, 8))
}
