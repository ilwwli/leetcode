package main

import (
	"fmt"
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	display := make(map[string][]int)
	menus := make(map[string]int) // name->index
	cuisine := make([]string, 0)

	for _, v := range orders {
		if _, ok := menus[v[2]]; !ok {
			menus[v[2]] = len(cuisine)
			cuisine = append(cuisine, v[2])
		}
	}
	sort.Sort(sort.StringSlice(cuisine))
	menus = make(map[string]int) // name->index
	for ind, v := range cuisine {
		menus[v] = ind
	}
	for _, v := range orders {
		table := v[1]
		if _, ok := display[table]; !ok {
			display[table] = make([]int, len(cuisine))
		}
		display[table][menus[v[2]]] += 1
	}
	return doDisplay(display, cuisine)
}

func doDisplay(display map[string][]int, cuisine []string) [][]string {
	ans := make([][]string, 0, len(display))
	for tablenum, v := range display {
		if len(v) < len(cuisine) {
			v = append(v, make([]int, len(cuisine)-len(v))...)
		}
		tmp := make([]string, 1, len(cuisine))
		tmp[0] = tablenum
		for _, v2 := range v {
			tmp = append(tmp, strconv.Itoa(v2))
		}
		ans = append(ans, tmp)
	}
	sort.Slice(ans, func(i, j int) bool {
		a1, _ := strconv.Atoi(ans[i][0])
		a2, _ := strconv.Atoi(ans[j][0])
		return a1 < a2
	})
	ans = append([][]string{append([]string{"Table"}, cuisine...)}, ans...)
	return ans
}

func main() {
	orders := [][]string{
		[]string{"James", "12", "Fried Chicken"},
		[]string{"Ratesh", "12", "Fried Chicken"},
		[]string{"Amadeus", "12", "Fried Chicken"},
		[]string{"Adam", "1", "Canadian Waffles"},
	}
	fmt.Println(displayTable(orders))
}
