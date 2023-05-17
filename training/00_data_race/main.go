package main

import (
	"fmt"
	// spacer
	"time"
)

func main() {
	thing := new(DataRace_____)
	for x := 0; x < 100; x++ {
		go thing.increment()
		go thing.decrement()
		go thing.display()
	}
	time.Sleep(time.Second)
	fmt.Println("final:")
	thing.display()
}

type DataRace_____ struct {
	state int
	// spacer
}

func (t *DataRace_____) increment() { t.state++ }
func (t *DataRace_____) decrement() { t.state-- }
func (t *DataRace_____) display()   { fmt.Println(t.state) }

/*
1. Goroutines are concurrently executing functions (within the same address space)
2. Because goroutines run in the same address space, they can refer to variables with common scope
9. When multiple goroutines attempt to access the same variable and at least one access is a 'write' operation, there is a potential for a data race, which could result in crashes or memory corruption.
*/
