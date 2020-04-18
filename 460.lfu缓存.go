package main

import "container/list"

/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU缓存
 */

// @lc code=start
type LFUCache struct {
	size int
	cap       int
	locater   map[int]*list.Element // key -> Element
	freqSaver map[int]*list.List    // freq -> List
	minfreq   int
}

type ElementStruct struct {
	key   int
	value int
	freq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		size: 0,
		cap:       capacity,
		locater:   make(map[int]*list.Element),
		freqSaver: make(map[int]*list.List),
		minfreq:   0,
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.locater[key]
	if !ok {
		return -1
	}
	oldFreq := node.Value.(ElementStruct).freq
	// update Node
	node.Value.(ElementStruct).freq = oldFreq + 1
	// remove from oldFreq
	this.freqSaver[oldFreq].Remove(node)
	if oldFreq == this.minfreq && this.freqSaver[oldFreq].Len() == 0 {
		this.minfreq++
	}
	newlist, ok := this.freqSaver[oldFreq+1] {

	}
	this.freqSaver
	this.ls.MoveToFront(node)
	return node.Value.(ElementStruct).value
}

func (this *LFUCache) Put(key int, value int) {
	node, ok := this.locater[key]
	if ok {
		node.Value = value
		return
	}
	if this.ls.Len() == this.cap {
		back := this.ls.Back()
		delete(this.locater, back.Value.(ElementStruct).key)
		this.ls.Remove(back)
	}
	ele := ElementStruct{
		key:   key,
		value: value,
	}

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
