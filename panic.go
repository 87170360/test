package main

import (
	"fmt"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("%v", e)
		}
	}()
	panic("test")
}
