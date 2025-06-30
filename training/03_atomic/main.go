package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	thing := &SafeAtomicPtr{state: new(atomic.Int32)}
	for range 100 {
		go thing.increment()
		go thing.decrement()
		go thing.display()
	}
	time.Sleep(time.Second)
	fmt.Println("final:")
	thing.display()
}

type SafeAtomicPtr struct {
	state *atomic.Int32
	// spacer
}

func (t *SafeAtomicPtr) increment() { t.state.Add(1) }
func (t *SafeAtomicPtr) decrement() { t.state.Add(-1) }
func (t *SafeAtomicPtr) display()   { fmt.Println(t.state.Load()) }

/*
14. Concurrency != Parallelism
15. A Go program that is limited to running a single goroutine at a time is said to be concurrent
16. A Go program that can run multiple goroutines at the same time (such as by multiple CPUs) is said to be parallel

runtime.GOMAXPROCS(runtime.NumCPU())

What 'HACK' do each of these programs share?

11. Go Slogan: "Do not communicate by sharing memory; instead, share memory by communicating."
*/
