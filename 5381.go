package main

import (
	"container/list"
	"fmt"
)

func processQueries(queries []int, m int) []int {
	l := list.New()
	for i := 1; i <= m; i++ {
		l.PushBack(i)
	}
	ans := []int{}
	for _, v := range queries {
		ind := 0
		for it := l.Front(); it != nil; it = it.Next() {
			if it.Value.(int) == v {
				ans = append(ans, ind)
				l.MoveToFront(it)
				break
			}
			ind++
		}
	}
	return ans
}

func main() {
	fmt.Println(processQueries([]int{}, 8))
}
