package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	ctx, cancel := context.WithCancel(context.Background())

	ch3 := asyncMerge[int](ctx, ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}

	cancel()
}

func asyncMerge[T any](ctx context.Context, chans ...chan T) chan T {
	result := make(chan T)
	wg := sync.WaitGroup{}

	for _, singleChan := range chans {
		wg.Add(1)
		singleChan := singleChan
		go func() {
			defer wg.Done()
			for val := range singleChan {
				result <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

// Это очевидно, но не очень здорово
func syncMerge[T any](chans ...chan T) chan T {
	l := 0
	for _, singleCh := range chans {
		l += len(singleCh)
	}

	res := make(chan T, l)

	for _, sighleCh := range chans {
		for Ch := range sighleCh {
			res <- Ch
		}
	}
	close(res)

	return res
}
