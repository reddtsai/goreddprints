package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// limit  max number of goroutine
	goroutineCount := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		goroutineCount <- i
		go func(num int) {
			fmt.Println(num, "s...")
			defer func() {
				fmt.Println(num, "e...")
				<-goroutineCount
				wg.Done()
			}()
			fmt.Println(num, "d...")
			time.Sleep(time.Second * 2)
		}(i)
	}
	wg.Wait()
}
