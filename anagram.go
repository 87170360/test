package main

import "fmt"

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	size := len(s1)
	for i := 0; i < size; i++ {
		if s1[i] != s2[size-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := "abc"
	s2 := "aba"
	fmt.Printf("%s %s %v\n", s1, s2, anagram(s1, s2))

	s1 = "abc"
	s2 = "cba"
	fmt.Printf("%s %s %v\n", s1, s2, anagram(s1, s2))

	s1 = "abcd"
	s2 = "dcba"
	fmt.Printf("%s %s %v\n", s1, s2, anagram(s1, s2))
}
