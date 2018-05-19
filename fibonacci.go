package main

import (
	"fmt"
)

func Fibonacci(num int) int {
	if num < 2 {
		return num
	}

	return Fibonacci(num-1) + Fibonacci(num-2)
}

func HofstadterFemale(n int) int {
	fmt.Println("HofstadterFemale n:", n)
	if n <= 0 {
		fmt.Println("HofstadterFemale return 1")
		return 1
	}
	fmt.Printf("%d - HofstadterMale(HofstadterFemale(%d - 1))\n", n, n)
	return n - HofstadterMale(HofstadterFemale(n-1))
}

func HofstadterMale(n int) int {
	fmt.Println("HofstadterMale n:", n)
	if n <= 0 {
		fmt.Println("HofstadterMale return 0")
		return 0
	}
	fmt.Printf("%d - HofstadterFemale(HofstadterMale(%d - 1))\n", n, n)
	return n - HofstadterFemale(HofstadterMale(n-1))
}

func main() {

	fmt.Println(HofstadterFemale(3))
}
