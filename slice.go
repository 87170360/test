package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	s = append(s, 1)
	fmt.Println(s)

	s1 := s[:0]
	fmt.Println(len(s1), cap(s1))
}
