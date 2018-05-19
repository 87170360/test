package main

import (
	"fmt"
)

func main() {

	found := false

	data := []int{1, 2, 3, 4, 5}
	data2 := []int{1, 2, 3, 4, 5}

FOUND:
	for _, v := range data {
		if v == 6 {
			for _, v1 := range data2 {
				if v1 == 5 {
					found = true
					break FOUND
				}
			}
		}
	}

	fmt.Println(found)
}
