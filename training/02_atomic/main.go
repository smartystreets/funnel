package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	thing := new(SafeAtomic___)
	for range 100 {
		go thing.increment()
		go thing.decrement()
		go thing.display()
	}
	time.Sleep(time.Second)
	fmt.Println("final:")
	thing.display()
}

type SafeAtomic___ struct {
	state int64
	// spacer
}

func (t *SafeAtomic___) increment() { atomic.AddInt64(&t.state, 1) }
func (t *SafeAtomic___) decrement() { atomic.AddInt64(&t.state, -1) }
func (t *SafeAtomic___) display()   { fmt.Println(atomic.LoadInt64(&t.state)) }

/*
3a. Another approach to safe concurrency is the use of the sync/atomic package for updating primitive values.
4. The main function of a Go program runs in a goroutine (running programs always have at least one goroutine)
5. When the 'main' goroutine exits, the entire process exits (along with any unfinished goroutines)

runtime.NumGoroutine()
*/
