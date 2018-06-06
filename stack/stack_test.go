//package stack_test
package main

import (
	"fmt"
	"test/stack"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

func TestStack(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	st := stack.Stack{}
	for _, v := range input {
		st.Append(v)
	}
	fmt.Println(st)
	fmt.Println(st.Top())
	fmt.Println(st.Pop())
	//fmt.Println(st.Remove(1))
	fmt.Println(st.RemoveFast(2))
}
