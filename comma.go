package main

import (
	"bytes"
	"fmt"
	"strings"
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

func commaFloat(s string) string {
	var prefix string
	var hasPrefix bool
	if hasPrefix = strings.HasPrefix(s, "-"); hasPrefix {
		prefix = s[:1]
		s = s[1:]
	}

	idx := strings.Index(s, ".")
	tmp1 := s[:idx]
	tmp2 := s[idx+1:]

	var buf bytes.Buffer
	if hasPrefix {
		buf.WriteString(prefix)
	}
	fmt.Fprintf(&buf, "%s.%s", comma(tmp1), comma(tmp2))
	return buf.String()
}

func main() {
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("1234"))

	fmt.Println(commaFloat("12300.45"))
	fmt.Println(commaFloat("12300.456"))
	fmt.Println(commaFloat("-12300.4567"))
}
