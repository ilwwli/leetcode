package main

import "fmt"

func destCity(paths [][]string) string {
	a := make(map[string]int)
	for _, v := range paths {
		a[v[0]] |= 2
		a[v[1]] |= 1
	}

	for key, value := range a {
		if value == 1 {
			return key
		}
	}
	return ""
}

func main() {
	fmt.Println(destCity([][]string{
		{"London", "New York"},
		{"New York", "Lima"},
		{"Lima", "Sao Paulo"},
	}))
}
