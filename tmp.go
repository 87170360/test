package main

import (
	"fmt"
	"math"
)

func func1() (bool, bool) {
	return true, true
}

func main() {
	fmt.Printf("max int32: %v\n", math.MaxInt32)
	fmt.Printf("max float32: %v\n", math.MaxFloat32)

	ok := false
	if ok, ok2 := func1(); ok {
		fmt.Println("cover")
		fmt.Println(ok2)
	}

	fmt.Println(ok)
}
