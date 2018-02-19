package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		m := rand.Perm(4)
		fmt.Println(m)

		list := rand.Perm(25)
		fmt.Println(list)

			for i, _ := range list {
				list[i]++
			}
			fmt.Println(list)
	*/

	slice := []int{1, 2, 3, 4}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(slice)) {
		val := slice[i]
		fmt.Println(val)
	}
}
