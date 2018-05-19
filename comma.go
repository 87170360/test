package main

import (
	"bytes"
	"fmt"
)

const (
	offset int = 3
)

func comma(s string) string {
	var buf bytes.Buffer
	len := len(s)
	for i := offset; ; i += offset {
		if i < len {
			fmt.Fprintf(&buf, "%s,", s[i-offset:i])
		} else {
			buf.WriteString(s[i-offset : len])
			break
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("1234"))
}
