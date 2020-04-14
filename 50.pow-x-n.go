package main

import "fmt"

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	ans := float64(1)
	tmp := x
	if n < 0 {
		tmp = 1.0 / tmp
		if n == int(-1<<31) {
			tmp *= tmp
			n >>= 1
		}
		n = -n
	}
	for i := 0; i < 31; i++ {
		if n>>i&1 == 1 {
			ans *= tmp
		}
		tmp *= tmp
	}
	return ans

}

// @lc code=end
func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}
