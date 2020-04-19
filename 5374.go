package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func getHappyString(n int, k int) string {
	tmp := 1
	for i := 0; i < n-1; i++ {
		tmp *= 2
	}
	k -= 1
	if k >= tmp*3 {
		return ""
	}
	ans := bytes.NewBufferString("")
	var last rune
	if k < tmp {
		last = 'a'
	} else if k < 2*tmp {
		last = 'b'
		k -= tmp
	} else {
		last = 'c'
		k -= 2 * tmp
	}
	ans.WriteRune(last)
	if n == 1 {
		return ans.String()
	}
	binary := strconv.FormatInt(int64(k), 2)
	padding := strings.Repeat("0", n-len(binary)-1)
	fmt.Println(padding + binary)
	for _, v := range padding + binary {
		lastb := last
		if v == '0' {
			if last == 'a' {
				last = 'b'
			} else {
				last = 'a'
			}
		} else if v == '1' {
			if last == 'c' {
				last = 'b'
			} else {
				last = 'c'
			}
		}
		fmt.Println(string(lastb), string(v), string(last))
		ans.WriteRune(last)
	}
	return ans.String()
}

func main() {
	//fmt.Println(getHappyString(3, 9))
	//fmt.Println(getHappyString(1, 3))
	//fmt.Println(getHappyString(2, 7))
	fmt.Println(getHappyString(10, 100))
}
