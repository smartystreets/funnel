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
