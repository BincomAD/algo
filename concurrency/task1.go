package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 20

	// Исправить этот код, чтобы он работал корректно
	//for i := 0; i <= counter; i++ {
	//	go func() {
	//		fmt.Println(i*i)
	//	}()
	//}
	//
	//time.Sleep(time.Second)

	wg := &sync.WaitGroup{}

	for i := 0; i <= counter; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(i)
	}

	wg.Wait()
}
