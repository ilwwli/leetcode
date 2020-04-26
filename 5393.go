package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}
	maxm := sum
	for i := 0; i < k; i++ {
		sum -= cardPoints[k-1-i]
		sum += cardPoints[len(cardPoints)-1-i]
		if sum > maxm {
			maxm = sum
		}
	}
	return maxm
}

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
	fmt.Println(maxScore([]int{2, 2, 2}, 2))
	fmt.Println(maxScore([]int{9, 7, 7, 9, 7, 7, 9}, 7))
	fmt.Println(maxScore([]int{1, 1000, 1}, 1))
	fmt.Println(maxScore([]int{1, 79, 80, 1, 1, 1, 200, 1}, 3))
}
