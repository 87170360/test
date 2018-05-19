package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	fmt.Println(len(s), cap(s))
	s = append(s, 10)
	fmt.Println(s)
}
