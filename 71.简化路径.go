package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=71 lang=golang
 *
 * [71] 简化路径
 */

// @lc code=start
func simplifyPath(path string) string {
	dirs := []string{}
	start := -1
	for ind, v := range path + "/" {
		if v == '/' {
			tmp := path[start+1 : ind]
			start = ind
			if tmp == "." || tmp == "" {
				continue
			}
			if tmp == ".." {
				if len(dirs) > 0 {
					dirs = dirs[:len(dirs)-1]
				}
				continue
			}
			dirs = append(dirs, tmp)
			continue
		}
	}
	return "/" + strings.Join(dirs, "/")
}

// @lc code=end
func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
