/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	M, N := len(matrix), len(matrix[0])
	dp := make([][]int, M)
	for j := 0; j < M; j++ {
		dp[j] = make([]int, N)
	}
	maxm := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == byte('0') {
				dp[i][j] = 0
			} else if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				minm := dp[i-1][j-1]
				if dp[i-1][j] < minm {
					minm = dp[i-1][j]
				}
				if dp[i][j-1] < minm {
					minm = dp[i][j-1]
				}
				dp[i][j] = minm + 1
			}
			if dp[i][j] > maxm {
				maxm = dp[i][j]
			}
		}
	}
	return maxm * maxm
}

// @lc code=end
