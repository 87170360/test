package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February"}
	q1 := months[:]
	q2 := months[:]
	q1[1] = "xx"
	fmt.Println(q1)
	fmt.Println(q2)
}
