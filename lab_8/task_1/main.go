package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m sync.Mutex
	v int
}

func (c *Counter) Inc() {
	c.m.Lock()
	c.v++
	c.m.Unlock()
}

func runOnce() int {
	c := Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	return c.v
}

func main() {
	for i := 1; i <= 10; i++ {
		result := runOnce()
		fmt.Printf("Результат #%d: %d\n", i, result)
	}
}
