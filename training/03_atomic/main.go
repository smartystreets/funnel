package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var state = new(atomic.Int64)

// spacer

func main() {
	for x := 0; x < 100; x++ {
		go increment()
		go decrement()
		go display()
	}
	time.Sleep(time.Second)
}

func increment() { state.Add(1) }
func decrement() { state.Add(-1) }
func display()   { fmt.Println(state.Load()) }
