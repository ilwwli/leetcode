package main

import (
	"fmt"
)

func numOfArrays(n int, m int, k int) int {
	return doNumOfArrays(n, m, k, 0, make(map[[4]int]int))
}

func doNumOfArrays(n, m, k int, max int, cache map[[4]int]int) int {
	if m-max < k || n < k {
		return 0
	}
	if n == 0 {
		return 1
	}
	key := [4]int{n, m, k, max}
	if v, ok := cache[key]; ok {
		return v
	}
	ans := 0
	// padding
	if max != 0 {
		tmp := doNumOfArrays(n-1, m, k, max, cache)
		for i := 0; i < max; i++ {
			ans += tmp
			ans %= 1000000007
		}
	}
	// ismax
	if k > 0 {
		for newmax := max + 1; newmax <= m-k+1; newmax++ {
			ans += doNumOfArrays(n-1, m, k-1, newmax, cache)
			ans %= 1000000007
		}
	}
	cache[key] = ans
	return ans
}

func main() {

	fmt.Println(numOfArrays(2, 3, 1))
	fmt.Println(numOfArrays(50, 100, 25))
}
