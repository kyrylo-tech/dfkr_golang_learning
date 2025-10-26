package main

import "fmt"

func main() {
	ch := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		fmt.Println("Надіслано:", i)
		ch <- i
	}

	ch <- 6

	for i := 0; i < 5; i++ {
		val := <-ch
		fmt.Println("Отримано:", val)
	}
}
