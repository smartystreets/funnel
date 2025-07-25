package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	thing := new(SafeLocking__)
	for range 100 {
		go thing.increment()
		go thing.decrement()
		go thing.display()
	}
	time.Sleep(time.Second)
	fmt.Println("final:")
	thing.display()
}

type SafeLocking__ struct {
	state int
	lock  sync.RWMutex
}

func (t *SafeLocking__) increment() { t.lock.Lock(); defer t.lock.Unlock(); t.state++ }
func (t *SafeLocking__) decrement() { t.lock.Lock(); defer t.lock.Unlock(); t.state-- }
func (t *SafeLocking__) display()   { t.lock.RLock(); defer t.lock.RUnlock(); fmt.Println(t.state) }

/*
3. One approach to safe concurrency is to use a 'mutex' (mutual exclusion lock)
*/
