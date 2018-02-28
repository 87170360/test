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
	m := map[*Point]string{&Point{1, 2, 3}: "a", &Point{4, 5, 6}: "b", &Point{7, 8, 9}: "c"}
	fmt.Println(m)
	m1 := map[string]*Point{"a": &Point{1, 2, 3}}
	fmt.Println(m1["a"])
}
