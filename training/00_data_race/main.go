package main

import (
	"fmt"
	"time"
)

// spacer
var state int

// spacer

func main() {
	for x := 0; x < 100; x++ {
		go increment()
		go decrement()
		go display()
	}
	time.Sleep(time.Second)
}

func increment() { state++ }
func decrement() { state-- }
func display()   { fmt.Println(state) }
