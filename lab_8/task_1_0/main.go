package main

import (
	"fmt"
	"sync"
)

var counter int

func runNoMutex() int {
	counter = 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()
	return counter
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Результат #%d: %d\n", i, runNoMutex())
	}
}
