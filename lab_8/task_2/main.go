package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CacheMap struct {
	m        sync.Mutex
	products map[string]int
}

func (c *CacheMap) Set(key string, value int) {
	c.m.Lock()
	c.products[key] = value
	c.m.Unlock()
}

func (c *CacheMap) Get(key string) int {
	c.m.Lock()
	defer c.m.Unlock()
	return c.products[key]
}

type RWCacheMap struct {
	m        sync.RWMutex
	products map[string]int
}

func (c *RWCacheMap) RWSet(key string, value int) {
	c.m.Lock()
	c.products[key] = value
	c.m.Unlock()
}

func (c *RWCacheMap) RWGet(key string) int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.products[key]
}

type Cache struct {
	DMap  *CacheMap
	RWMap *RWCacheMap
}

func runBenchmark(doRead func(), doWrite func(), numGoroutines int, readPercent float64) time.Duration {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if rand.Float64() < readPercent {
				doRead()
			} else {
				doWrite()
			}
		}()
	}

	wg.Wait()
	return time.Since(start)
}

func fill(m map[string]int) {
	for i := 0; i < 100; i++ {
		m[fmt.Sprintf("item_%d", i)] = i
	}
}

func main() {
	c := Cache{
		DMap:  &CacheMap{products: map[string]int{}},
		RWMap: &RWCacheMap{products: map[string]int{}},
	}

	fill(c.DMap.products)
	fill(c.RWMap.products)

	fmt.Println("== Benchmark ==")

	tests := []struct {
		name        string
		readPercent float64
	}{
		{"90% читання", 0.9},
		{"50% / 50%", 0.5},
		{"10% читання", 0.1},
	}

	for _, t := range tests {
		fmt.Println("\nСценарій:", t.name)

		d1 := runBenchmark(
			func() { c.DMap.Get("item_1") },
			func() { c.DMap.Set("item_1", rand.Intn(1000)) },
			100000,
			t.readPercent,
		)

		d2 := runBenchmark(
			func() { c.RWMap.RWGet("item_1") },
			func() { c.RWMap.RWSet("item_1", rand.Intn(1000)) },
			100000,
			t.readPercent,
		)

		fmt.Printf("sync.Mutex:   %v\n", d1)
		fmt.Printf("sync.RWMutex: %v\n", d2)
		ratio := float64(d1.Milliseconds()) / float64(d2.Milliseconds())
		fmt.Printf(">> RWMutex швидший у %.2f разів\n", ratio)
	}
}
