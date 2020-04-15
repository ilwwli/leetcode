/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int)  {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	M, N := len(matrix), len(matrix[0])
	flagR,flagC := false,false
	// #1
	for i := 0;i<M;i++ {
		for j:= 0;j<N;j++ {
			if matrix[i][j] == 0 {
				if i== 0 && j == 0 {
					flagR = true
					flagC = true
					continue
				}
				if i == 0 {
					flagR = true
				} else {
					matrix[0][j] = 0
				}
				if j == 0 {
					flagC = true
				} else {
					matrix[i][0] = 0
				}
			}
		}
	}
	// #2
	for i := 1;i<M;i++ {
		for j:= 1;j<N;j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0{
				matrix[i][j] = 0
			}
		}
	}
	if flagR {
		for j:=0;j<N;j++ {
			matrix[0][j] = 0
		}
	}
	if flagC {
		for i:=0;i<M;i++ {
			matrix[i][0] = 0
		}
	}
	return
}
// @lc code=end

