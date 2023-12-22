package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	workerPool()
}

func workerPool() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inCh, outCh := make(chan int, 5), make(chan int, 5)
	wg := &sync.WaitGroup{}

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, inCh, outCh)
		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			inCh <- i
		}
		close(inCh)
	}()

	go func() {
		wg.Wait()
		close(outCh)
	}()

	for i := range outCh {
		fmt.Println(i)
	}
}

func worker(ctx context.Context, inNumbers <-chan int, outNumbers chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-inNumbers:
			if !ok {
				return
			}
			outNumbers <- val * val
		}
	}

}
