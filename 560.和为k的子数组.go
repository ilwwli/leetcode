/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	sumprefix := make(map[int]int, len(nums))
	sumprefix[0]++
	summ, ans := 0, 0
	for _, v := range nums {
		summ += v
		ans += sumprefix[summ-k]
		sumprefix[summ]++
	}
	return ans
}

// @lc code=end
