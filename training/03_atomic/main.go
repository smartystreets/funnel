package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	thing := &SafeAtomicPtr{state: new(atomic.Int32)}
	for x := 0; x < 100; x++ {
		go thing.increment()
		go thing.decrement()
		go thing.display()
	}
	time.Sleep(time.Second)
}

type SafeAtomicPtr struct {
	state *atomic.Int32
	// spacer
}

func (t *SafeAtomicPtr) increment() { t.state.Add(1) }
func (t *SafeAtomicPtr) decrement() { t.state.Add(-1) }
func (t *SafeAtomicPtr) display()   { fmt.Println(t.state.Load()) }
