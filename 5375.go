package main

import (
	"fmt"
	"strconv"
)

func numberOfArrays(s string, k int) int {
	n := strconv.Itoa(k)
	N := len(n)

	dp := []int{1}
	for i := 1; i <= len(s); i++ {
		tmp := 0
		for j := i - 1; i-j <= N && j >= 0; j-- {
			if s[j] == byte('0') {
				continue
			}
			if num, _ := strconv.Atoi(s[j:i]); num <= k {
				tmp += dp[j]
				tmp %= 1000000007
			}
		}
		dp = append(dp, tmp)
	}
	return dp[len(dp)-1]
}

func main() {

	fmt.Println(numberOfArrays("1000", 10))
	fmt.Println(numberOfArrays("1317", 2000))
	fmt.Println(numberOfArrays("1234567890", 90))
}
