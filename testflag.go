package main

import (
	"flag"
	"fmt"
)

//go run testflag.go -f1=1 -f2=t aa bb cc
func main() {
	flag.Int("f1", 123456, "help message for f1")
	flag.Bool("f2", true, "help message for f2")
	flag.String("f3", "f3 message", "help message for f3")
	flag.Bool("h", false, "help message for h")
	flag.Parse()
	//flag.Usage()
	//fmt.Println(flag.CommandLine)
	//fmt.Println(flag.Args())
	//fmt.Println(flag.Arg(1))
	//fmt.Println(flag.NArg())
	//fmt.Println(flag.NFlag())
	flag.Visit(func(fg *flag.Flag) {
		fmt.Println(fg.Name, fg.Value)
	})
	//fmt.Println(*f1)
	//fmt.Println(flag.Lookup("f1").Value)
	//fmt.Println(flag.Lookup("f2").Value)
	//fmt.Println(flag.Lookup("f2").DefValue)
	//fmt.Println(flag.Lookup("f3").Value)
	//fmt.Println(flag.Lookup("f4"))
}
