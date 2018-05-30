package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 { //退出条件i < j
		s[i], s[j] = s[j], s[i] //换值操作
	}
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	//reverse(a)
	reverse(a[:])
	fmt.Println(a)
}
