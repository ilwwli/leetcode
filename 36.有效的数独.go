package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := board[i]
		column := []byte{}
		for j := 0; j < 9; j++ {
			column = append(column, board[j][i])
		}
		square := []byte{}

		startR, startC := i/3*3, (i-i/3*3)*3
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				square = append(square, board[startR+j][startC+k])
			}
		}
		if !(isValid(row) && isValid(column) && isValid(square)) {
			return false
		}
	}
	return true
}

func isValid(s []byte) bool {
	exist := map[byte]bool{}
	for _, v := range s {
		if v == '.' {
			continue
		}
		if _, ok := exist[v]; ok {
			return false
		}
		exist[v] = true
	}
	return true
}

// @lc code=end
func main() {
	block := [][]byte{
		[]byte{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		[]byte{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		[]byte{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		[]byte{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		[]byte{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	fmt.Println(isValidSudoku(block))
}
