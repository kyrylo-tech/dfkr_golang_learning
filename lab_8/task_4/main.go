package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type MutexCounter struct {
	m sync.Mutex
	v int
}

func (c *MutexCounter) Inc() {
	c.m.Lock()
	c.v++
	c.m.Unlock()
}

func runMutex(counter *MutexCounter, users int, ops int) time.Duration {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < users; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < ops; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

type AtomicCounter struct {
	v int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.v, 1)
}

func runAtomic(counter *AtomicCounter, users int, ops int) time.Duration {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < users; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < ops; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

func main() {
	users := 1000
	ops := 10000
	expected := users * ops

	fmt.Println("Початок бенчмарку…")
	fmt.Printf("Очікуване значення: %d\n", expected)

	mc := &MutexCounter{}
	t1 := runMutex(mc, users, ops)
	fmt.Printf("\nsync.Mutex: Час: %v, Результат: %d\n", t1, mc.v)

	ac := &AtomicCounter{}
	t2 := runAtomic(ac, users, ops)
	fmt.Printf("sync.Atomic: Час: %v, Результат: %d\n", t2, ac.v)

	fmt.Printf("\n>> Атомарні операції швидші у %.2f разів\n",
		float64(t1.Milliseconds())/float64(t2.Milliseconds()))
}
