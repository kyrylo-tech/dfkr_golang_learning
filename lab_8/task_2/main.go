package main

import (
	"sync"
	"time"
)



type CacheMap struct {
	m sync.Mutex
	products map[string]int
}

func (c *CacheMap) Set(key string, value int) {
	return
}

func (c *CacheMap) Get(key string) {
	return
}



type RWCacheMap struct {
	m sync.RWMutex
	products map[string]int
}

func (c *RWCacheMap) RWSet(key string, value int) {
	return
}

func (c *RWCacheMap) RWGet(key string) {

}



type Cache struct {
	DMap CacheMap
	RWMap RWCacheMap
}


func runBenchmark(c Cache, numGos int, readPercent float64) time.Duration {
	return time.Duration
}

func printBenchmarkRes(clock time.Duration) {

}

func main() {
	c := Cache{
		DMap: CacheMap{},
		RWMap: RWCacheMap{},
	}

	runBenchmark(c, 100, .9)
	runBenchmark(c, 100, .5)
	runBenchmark(c, 100, .1)
}