package main

import "fmt"

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
type DisjointSetsNode struct {
	Val    int
	Parent *DisjointSetsNode
	Rank   int
}

type DisjointSets struct {
	Nodes map[int]*DisjointSetsNode
}

func NewDisjointSets() *DisjointSets {
	return &DisjointSets{make(map[int]*DisjointSetsNode)}
}

func (s *DisjointSets) Find(x int) *DisjointSetsNode {
	return s.Nodes[x]
}

func (s *DisjointSets) MakeSet(x int) *DisjointSetsNode {
	node := &DisjointSetsNode{x, nil, 1}
	node.Parent = node
	s.Nodes[x] = node
	return node
}

func (s *DisjointSets) FindSet(x int) *DisjointSetsNode {
	node, ok := s.Nodes[x]
	if !ok {
		return nil
	}
	return s.findSet(node)
}

func (s *DisjointSets) Union(x, y int) {
	if s.Nodes[x] == nil || s.Nodes[y] == nil {
		return
	}
	px, py := s.FindSet(x), s.FindSet(y)
	if px == py {
		return
	}
	if px.Rank > py.Rank {
		py.Parent = px
		px.Rank += py.Rank
	} else {
		px.Parent = py
		py.Rank += px.Rank
	}
}

func (s *DisjointSets) findSet(x *DisjointSetsNode) *DisjointSetsNode {
	if x.Parent != x {
		x.Parent = s.findSet(x.Parent)
	}
	return x.Parent
}

func longestConsecutive(nums []int) int {
	sets := NewDisjointSets()
	for _, v := range nums {
		if sets.Find(v) != nil {
			continue
		}
		sets.MakeSet(v)
		sets.Union(v-1, v)
		sets.Union(v, v+1)
	}
	maxm := 0
	for _, v := range sets.Nodes {
		if v.Rank > maxm {
			maxm = v.Rank
		}
	}
	return maxm
}

// @lc code=end
func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
