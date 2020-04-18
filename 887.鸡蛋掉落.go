package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=887 lang=golang
 *
 * [887] 鸡蛋掉落
 */

// @lc code=start
func superEggDrop(K int, N int) int {
	N++
	for base := 0; ; base++ {
		maxN := 0
		for i := 0; i <= K && i <= base; i++ {
			maxN += comb(i, base)
		}
		if maxN >= N {
			return base
		}
	}
}

// 计算组合数C^m_n
func comb(m, n int) int {
	if m > n {
		m = n
	}
	if m > n/2 {
		m = n - m
	}
	ans := 1
	for i := n - m + 1; i <= n; i++ {
		ans *= i
	}
	for i := 2; i <= m; i++ {
		ans /= i
	}
	return ans
}

// @lc code=end

func main() {
	fmt.Println(superEggDrop(1, 1))
	fmt.Println(superEggDrop(1, 2))
	fmt.Println(superEggDrop(2, 3))
	fmt.Println(superEggDrop(3, 14))
	fmt.Println(superEggDrop(4, 10000))
	fmt.Println(superEggDrop(100, 10000))
}
