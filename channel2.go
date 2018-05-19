package main

import (
	"fmt"
)

func expensiveComputation(data int, answer chan int, done chan bool) {
	finished := false
	result := 1
	for !finished {
		answer <- result
		finished = true
	}
	done <- true
}

func main() {
	const allDone = 2
	doneCount := 0
	answer1 := make(chan int)
	answer2 := make(chan int)
	done1 := make(chan bool)
	done2 := make(chan bool)
	defer func() {
		close(answer1)
		close(answer2)
		close(done1)
		close(done2)
	}()

	go expensiveComputation(1, answer1, done1)
	go expensiveComputation(2, answer2, done2)

	for doneCount != allDone {
		var which string
		var result int
		select {
		case result = <-answer1:
			which = "chan 1 result"
			fmt.Println("result", result)
		case result = <-answer2:
			which = "chan 2 result"
			fmt.Println("result", result)
		case <-done1:
			doneCount++
			which = "chan 1 done"
		case <-done2:
			doneCount++
			which = "chan 2 done"
		}

		fmt.Printf("[%s]\n", which)
	}
	fmt.Println()

}
