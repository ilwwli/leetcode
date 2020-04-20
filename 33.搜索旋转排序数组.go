package main

import "fmt"

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // front part ordered
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
	return -1
}

// @lc code=end
func main() {
	fmt.Println(search([]int{
		4, 5, 6, 7, 0, 1, 2,
	}, 0))
	fmt.Println(search([]int{
		4, 5, 6, 7, 0, 1, 2,
	}, 3))
	fmt.Println(search([]int{
		7, 8, 1, 2, 3, 4, 5, 6,
	}, 2))
}
