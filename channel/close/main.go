package main

import (
	"fmt"
	"sync"
)

type sample struct {
	ch   chan int
	once sync.Once
}

func main() {
	sample := &sample{
		ch: make(chan int),
	}
	sample.closeChannelOnce()
	sample.closeChannelOnce()
	sample.handleClosePanic()
}

func (s *sample) closeChannelOnce() {
	// 必免重複關閉通道例外
	s.once.Do(func() {
		fmt.Println("close channel once")
		close(s.ch)
	})
}

func (s *sample) handleClosePanic() {
	// 恢復重複關閉通道例外
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover -", r)
		}
	}()
	close(s.ch)
}
