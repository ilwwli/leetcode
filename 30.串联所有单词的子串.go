package main

import (
	"fmt"
	"sync"
)

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	if s == "" || len(words) == 0 {
		return []int{}
	}
	wordss := make(map[string]int)
	for _, word := range words {
		wordss[word]++
	}
	l := len(words[0])
	wg := sync.WaitGroup{}
	ans := make(chan int)
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func(i int) {
			finder(s[i:], wordss, l, ans, i)
			wg.Done()
		}(i)

	}
	go func() {
		wg.Wait()
		close(ans)
	}()
	list := make([]int, 0)
	for v := range ans {
		list = append(list, v)
	}
	return list
}

func finder(s string, words map[string]int, l int, ans chan<- int, margin int) {
	tmpmap := make(map[string]int)
	start := 0
	for i := 0; i < len(s)/l; i++ {
		word := s[l*i : l*(i+1)]
		if words[word] == 0 {
			for key := range tmpmap {
				delete(tmpmap, key)
			}
			start = i + 1
			continue
		}
		if words[word] <= tmpmap[word] {
			for ; start < i; start++ {
				dword := s[l*start : l*(start+1)]
				if word == dword {
					start++
					break
				}
				tmpmap[dword]--
			}
			continue
		}
		tmpmap[word]++
		//fmt.Println(start, tmpmap)
		if equal(words, tmpmap) {
			ans <- start*l + margin
			dword := s[l*start : l*(start+1)]
			tmpmap[dword]--
			start++
		}
	}
}

func equal(a, b map[string]int) bool {
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

// @lc code=end

func main() {
	fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
}
