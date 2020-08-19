package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go work1(wg)
	work2(wg)
	wg.Wait()
}

func work1(wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(time.Second * 4):
		log.Println("work1 finish")
	}
}

func work2(wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(time.Second * 2):
		log.Println("work2 finish")
	}
}
