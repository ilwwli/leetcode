package main

import "fmt"

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	doSolveSudoku(board, 0)
}

func doSolveSudoku(board [][]byte, r int) bool {
	for i := r; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			for _, v := range validnums(board, i, j) {
				board[i][j] = v
				if doSolveSudoku(board, i) {
					return true
				}
			}
			board[i][j] = '.'
			return false
		}
	}
	return true
}

func validnums(board [][]byte, row, column int) []byte {
	option := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
	for i := 0; i < 9; i++ {
		if board[row][i] != '.' {
			option[board[row][i]-'1'] = '.'
		}
		if board[i][column] != '.' {
			option[board[i][column]-'1'] = '.'
		}
	}
	row = row / 3 * 3
	column = column / 3 * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[row+i][column+j] != '.' {
				option[board[row+i][column+j]-'1'] = '.'
			}
		}
	}
	ans := make([]byte, 0)
	for _, v := range option {
		if v != '.' {
			ans = append(ans, v)
		}
	}
	return ans
}

// @lc code=end

func main() {
	block := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(block)
	fmt.Println(block)
}
