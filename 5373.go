package main

import "fmt"

func findMinFibonacciNumbers(k int) int {
	fib := []int{1, 1}
	for fib[len(fib)-1] < k {
		fib = append(fib, fib[len(fib)-2]+fib[len(fib)-1])
	}
	times := 0
	for i := len(fib) - 1; i >= 0; i-- {
		if fib[i] <= k {
			k -= fib[i]
			times++
		}
		if k == 0 {
			break
		}
	}
	return times

}

func main() {
	fmt.Println(findMinFibonacciNumbers(7))
	fmt.Println(findMinFibonacciNumbers(10))
}
