package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	thing := new(SafeAtomic___)
	for x := 0; x < 100; x++ {
		go thing.increment()
		go thing.decrement()
		go thing.display()
	}
	time.Sleep(time.Second)
}

type SafeAtomic___ struct {
	state int64
	// spacer
}

func (t *SafeAtomic___) increment() { atomic.AddInt64(&t.state, 1) }
func (t *SafeAtomic___) decrement() { atomic.AddInt64(&t.state, -1) }
func (t *SafeAtomic___) display()   { fmt.Println(atomic.LoadInt64(&t.state)) }
