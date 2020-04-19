package main

/*
 * @lc app=leetcode.cn id=466 lang=golang
 *
 * [466] 统计重复个数
 */

// @lc code=start
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	dp := make(map[int][2]int) // match_times -> ind1, times
	matchTimeCache := make([]int, 0)
	s2 = strings.Repeat(s2, n2)
	ind1, ind2 := 0, 0
	times := 1
	match_times := 0
	for {
		if s1[ind1] == s2[ind2] {			
			ind2++
			if ind2 == len(s2) {
				dp[match_times] = ind1
				matchTimeCache = append(matchTimeCache, times)
			}
			ind1++
			if ind1 == len(s1) {
				times++
				ind1 = 0
			}			
		}
	}
}

// @lc code=end
func main() {

}
-+
