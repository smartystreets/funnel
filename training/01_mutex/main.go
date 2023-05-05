package main

import (
	"fmt"
	"sync"
	"time"
)

var state int

var lock sync.RWMutex

func main() {
	for x := 0; x < 100; x++ {
		go increment()
		go decrement()
		go display()
	}
	time.Sleep(time.Second)
}

func increment() { lock.Lock(); defer lock.Unlock(); state++ }
func decrement() { lock.Lock(); defer lock.Unlock(); state-- }
func display()   { lock.RLock(); defer lock.RUnlock(); fmt.Println(state) }
