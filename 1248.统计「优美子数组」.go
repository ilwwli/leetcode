package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1248 lang=golang
 *
 * [1248] 统计「优美子数组」
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	left, right := -1, 0
	odd := 0
	evenleft := 1
	cnt := 0
	for right < len(nums) {
		evenleft = 1
		for left++; left < len(nums) && nums[left]%2 != 1; left++ {
			evenleft++
		}
		for ; right < len(nums) && odd <= k; right++ {
			if nums[right]%2 == 1 {
				odd++
			}
			if odd == k {
				cnt += evenleft
			}
		}
		odd--
	}
	if nums[right-1]%2 == 1 {
		evenleft = 1
		for left++; left < len(nums) && nums[left]%2 != 1; left++ {
			evenleft++
		}
		cnt += evenleft
	}
	return cnt
}

// @lc code=end
func main() {
	// fmt.Println(numberOfSubarrays([]int{
	// 	1, 1, 2, 1, 1,
	// }, 3))
	// fmt.Println(numberOfSubarrays([]int{
	// 	2, 4, 6,
	// }, 1))
	fmt.Println(numberOfSubarrays([]int{
		2, 2, 2, 1, 2, 2, 1, 2, 2, 2, 1,
	}, 2))
}
