package main

import (
	"fmt"
	"sync"
)

var links = []string{....}

func main() {
	counter := 0
	inCh := make(chan string, len(links))

	for i := range links{
		inCh <- links[i]
	}

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range inCh {
				if checkUrl(i) == nil{
					mu.Lock()
					counter += 1
					mu.Unlock()
				}
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func checkUrl(url string) (err error) {
	return nil
}
