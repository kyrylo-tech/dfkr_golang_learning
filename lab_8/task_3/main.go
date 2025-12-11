package main

import (
	"fmt"
	"sync"
)

type Config struct {
	host  string
	port  int
	debug bool
}

var (
	config *Config
	once   sync.Once
)

func initConfig() {
	fmt.Println("Конфігурацію завантажено...")
	config = &Config{
		host:  "localhost",
		port:  8080,
		debug: true,
	}
}

func GetConfig() *Config {
	once.Do(initConfig)
	return config
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			c := GetConfig()
			fmt.Printf("Goroutine #%d: %s:%d (debug=%v)\n", id, c.host, c.port, c.debug)
		}(i)
	}

	wg.Wait()
}
