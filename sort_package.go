package main

import (
	"fmt"
	"sort"
)

type Data struct {
	data []int
}

func (d Data) Len() int {
	return len(d.data)
}

func (d Data) Less(i, j int) bool {
	return d.data[i] < d.data[j]
}

func (d Data) Swap(i, j int) {
	tmp := i
	i = j
	j = tmp
}

func main() {
	fmt.Println("Hello,World!")
	d := []int{1, 2, 3}
	d1 := Data{d}
	r := sort.IsSorted(d1)
	fmt.Println(r)
}
