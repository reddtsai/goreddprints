package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 8)

	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("channel receive")
			default:
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer func() {
			ticker.Stop()
		}()
		for {
			select {
			case <-ticker.C:
				ch <- 1
			}
		}
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("done")
	}
}
