package main

import "fmt"

func waysToChange(n int) int {
	coins := []int{1, 2, 5}
	dp := make([]int, n/5+1)
	for i := 0; i <= n/5; i++ {
		dp[i] = 1
	}
	for _, coin := range coins {
		for i := coin; i <= n/5; i++ {
			dp[i] += dp[i-coin]
			dp[i] %= 1000000007
		}
	}
	return dp[n/5]
}

func main() {
	fmt.Println(waysToChange(500))
}
