package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	maxm := 0
	left, right := 0, len(height)-1
	for left < right {
		tmp := min(height[left], height[right]) * (right - left)
		if tmp > maxm {
			maxm = tmp
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxm
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// @lc code=end
func main() {

}
