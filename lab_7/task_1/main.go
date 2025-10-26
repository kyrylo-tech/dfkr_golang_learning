package main

import (
	"fmt"
	"time"
)

func sayHello(id int) {
	fmt.Printf("Привіт з потоку №%d\n", id)
}

func main() {
	go sayHello(1)
	go sayHello(2)
	go sayHello(3)

	time.Sleep(300 * time.Millisecond)
}
