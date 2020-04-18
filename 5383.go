package main

import "fmt"

func numOfWays(n int) int {
	cases := []string{
		"ryr", "yry", "gry",
		"ryg", "yrg", "grg",
		"rgr", "ygr", "gyr",
		"rgy", "ygy", "gyg",
	}
	caseIndexs := make([][]int, 12)
	for i := 0; i < len(cases); i++ {
		caseIndexs[i] = []int{}
		for j := 0; j < len(cases); j++ {
			if cases[i][0] != cases[j][0] &&
				cases[i][1] != cases[j][1] &&
				cases[i][2] != cases[j][2] {
				caseIndexs[i] = append(caseIndexs[i], j)
			}
		}

	}
	ans := []int{
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
		1, 1, 1,
	}
	for times := 1; times < n; times++ {
		tmp := make([]int, 12)
		for i := 0; i < len(cases); i++ {
			for _, v := range caseIndexs[i] {
				tmp[i] += ans[v]
				tmp[i] %= 1000000007
			}
		}
		ans = tmp
	}
	result := 0
	for _, v := range ans {
		result += v
		result %= 1000000007
	}
	return result
}

func main() {
	fmt.Println(numOfWays(5000))
}
