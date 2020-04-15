package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
type Element struct {
	i, j   int
	length int
}

type EleListNode struct {
	*Element
	Next *EleListNode
}

type EleQueue struct {
	head, tail *EleListNode
	len        int
}

func (q EleQueue) Len() int {
	return q.len
}

func (q *EleQueue) PushBack(ele *Element) {
	node := &EleListNode{ele, nil}
	if q.tail == nil {
		q.tail = node
		q.head = q.tail
	} else {
		q.tail.Next = node
		q.tail = q.tail.Next
	}
	q.len++
}

func (q *EleQueue) PopFront() *Element {
	if q.head == nil {
		return nil
	}
	front := q.head.Element
	if q.tail == q.head {
		q.tail = q.head.Next
	}
	q.head = q.head.Next
	q.len--
	return front
}

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	M, N := len(matrix), len(matrix[0])
	ans := make([][]int, 0, M)
	seen := make([][]bool, 0, M)
	for i := 0; i < M; i++ {
		tmp := make([]int, N, N)
		for j := 0; j < N; j++ {
			tmp[j] = N + M
		}
		ans = append(ans, tmp)
		seen = append(seen, make([]bool, N, N))
	}
	elelist := &EleQueue{}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == 0 {
				seen[i][j] = true
				elelist.PushBack(&Element{
					i:      i,
					j:      j,
					length: 0,
				})
			}
		}
	}
	for elelist.Len() > 0 {
		ele := elelist.PopFront()
		i, j, length := ele.i, ele.j, ele.length
		if length >= ans[i][j] {
			continue
		}
		ans[i][j] = length
		if i > 0 && !seen[i-1][j] {
			seen[i-1][j] = true
			elelist.PushBack(&Element{
				i:      i - 1,
				j:      j,
				length: length + 1,
			})
		}
		if i < M-1 && !seen[i+1][j] {
			seen[i+1][j] = true
			elelist.PushBack(&Element{
				i:      i + 1,
				j:      j,
				length: length + 1,
			})
		}
		if j > 0 && !seen[i][j-1] {
			seen[i][j-1] = true
			elelist.PushBack(&Element{
				i:      i,
				j:      j - 1,
				length: length + 1,
			})
		}
		if j < N-1 && !seen[i][j+1] {
			seen[i][j+1] = true
			elelist.PushBack(&Element{
				i:      i,
				j:      j + 1,
				length: length + 1,
			})
		}
	}
	return ans
}

// @lc code=end
func main() {
	fmt.Println(updateMatrix([][]int{
		[]int{0},
		[]int{0},
		[]int{0},
		[]int{0},
		[]int{0},
	}))
}
