package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var state int64

func main() {
	for x := 0; x < 100; x++ {
		go increment()
		go decrement()
		go display()
	}
	time.Sleep(time.Second)
}

func increment() {
	atomic.AddInt64(&state, 1)
}
func decrement() {
	atomic.AddInt64(&state, -1)
}
func display() {
	fmt.Println(state)
}
