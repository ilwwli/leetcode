package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=365 lang=golang
 *
 * [365] 水壶问题
 */

// @lc code=start
func canMeasureWater(x int, y int, z int) bool {
	if z < 0 {
		return false
	}
	if x <= 0 || y <= 0 {
		return x == z || y == z
	}
	return z%gcd(x, y) == 0 && z <= x+y
}

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

// @lc code=end

func main() {
	fmt.Println(canMeasureWater(3, 5, 4))
}
