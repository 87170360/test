package main

import (
	"fmt"
)

type Product struct {
	name  string
	price float64
}

func (p Product) String() string {
	return fmt.Sprintf("%s(%.2f)", p.name, p.price)
}

func main() {
	fmt.Println("Hello,World!")
	p := &Product{"peter", 123}
	fmt.Printf("%v", p)
}
