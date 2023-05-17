package main

import (
	"fmt"
	"time"

	"github.com/smartystreets/funnel/training/internet"
)

func main() {
	started := time.Now()
	defer func() { fmt.Println(time.Since(started)) }()

	for a, address := range internet.Addresses {
		fmt.Println(a, internet.Scrape(address))
	}
}

/*
https://go.dev/talks/2012/waza.slide#12
*/
