package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	m[1] = 2
	m[2] = 3
	delete(m, 1)
	delete(m, 5)
	fmt.Println(m)
}
