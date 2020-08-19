package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	go func() {
		select {
		case <-time.After(time.Second * 3):
			cancel()
		}
	}()

	loop(ctx)
}

func loop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		}
	}
}
