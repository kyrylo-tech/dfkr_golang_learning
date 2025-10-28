package main

import (
	"fmt"
	"time"
	"sync"
)

type counter struct {
	m sync.Mutex
	v int
}

func (c *counter) incr() {
	c.m.Lock()
	c.v++
	c.m.Unlock()
}

func main() {
	c := counter{v: 0}

	for i := 0; i < 1000; i++ {
		go c.incr()
	}

	time.Sleep(time.Second)
	fmt.Printf("\nSafe counter: %d\n\n", c.v)
}