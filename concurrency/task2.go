package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 1000
	storage := make(map[int]int, writes)
	mu := sync.Mutex{}

	wg := sync.WaitGroup{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()
			mu.Lock()
			storage[i] = i
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}
