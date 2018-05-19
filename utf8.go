package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // 13 所占字节总数
	fmt.Println(utf8.RuneCountInString(s)) // 9 unicode 编码下的个数 'Hello, ' 7个 + 汉字2个

	//打印每一个unicode编码字符所在位置
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	//更简洁的方式
	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
		//打印的字节位置，可以看出所占长度，汉字占3个字节
	}
}
