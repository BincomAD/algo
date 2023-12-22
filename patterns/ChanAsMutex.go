package main

import "sync"

func main() {
	chAsMutex()
}

func chAsMutex() {
	var count int
	ch := make(chan struct{}, 1)
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			ch <- struct{}{}

			count++

			<-ch
			wg.Done()
		}()
	}

	wg.Wait()
	print(count)
}
