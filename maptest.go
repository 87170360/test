package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
	Z int
}

func (p *Point) String() string {
	return fmt.Sprintf("x:%d,y:%d,z:%d", p.X, p.Y, p.Z)
}

func main() {
	m := make(map[int32]int32)
	m[1] = 1
	m[2] = 2
	m[3] = 3
	m[4] = 4
	fmt.Println(m)
}
