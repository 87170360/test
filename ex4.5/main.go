package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	fmt.Println(eliminateAdjacent([]string{"a", "a", "b", "c", "b", "d", "d", "d", "e", "f"}))
}

func eliminateAdjacent(strs []string) []string {
	idx := 0
	str := ""
	for _, v := range strs {
		if str == v {
			continue
		}
		strs[idx] = v
		str = v
		idx++
	}
	return strs[:idx]
}
