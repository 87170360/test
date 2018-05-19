package main

import (
	"fmt"
	"time"
)

func getTimeByHour(hour int) time.Time {
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), hour, 0, 0, 0, time.Local)
	return t
}

func main() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	t1 := time.Now()
	fmt.Printf("Go now at %s\n", t1.Local())
	fmt.Printf("Go now at %s\n", getTimeByHour(12).Local())
}
