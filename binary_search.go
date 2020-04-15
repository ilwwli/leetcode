package main

import "fmt"

func binary_search_leftmost(A []int, n int, target int) int {
	L := 0
	R := n
	for L < R {
		m := (L + R) / 2
		if A[m] < target {
			L = m + 1
		} else {
			R = m
		}
	}

	return L
}

func binary_search_leftmost2(A []int, n int, target int) int {
	L := 0
	R := n - 1
	for L < R {
		m := (L + R) / 2
		if A[m] < target {
			L = m + 1
		} else {
			R = m
		}
	}

	return L
}

func binary_search_rightmost(A []int, n int, target int) int {
	L := 0
	R := n
	for L < R {
		m := (L + R) / 2
		if A[m] > target {
			R = m
		} else {
			L = m + 1
		}

	}

	return R - 1
}

func binary_search_rightmost2(A []int, n int, target int) int {
	L := 0
	R := n - 1
	for L < R {
		m := (L + R) / 2
		if A[m] > target {
			R = m
		} else {
			L = m + 1
		}

	}

	return R
}

func testbench(A []int, n int, target int) {
	fmt.Println(A)
	fmt.Println(
		binary_search_leftmost(A, n, target),
		binary_search_leftmost2(A, n, target),
		binary_search_rightmost(A, n, target),
		binary_search_rightmost2(A, n, target),
	)
}

func main() {
	testcases := [][]int{
		[]int{},
		[]int{0},
		[]int{1},
		[]int{2},
		[]int{0, 0},
		[]int{2, 2},
		[]int{0, 1},
		[]int{1, 2},
		[]int{0, 2},
		[]int{1, 1},
	}
	for _, c := range testcases {
		testbench(c, len(c), 1)
	}
}
