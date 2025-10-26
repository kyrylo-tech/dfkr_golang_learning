package main

import "fmt"

func sendNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Надсилається:", i)
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go sendNumbers(ch)

	for num := range ch {
		fmt.Println("Отримано:", num)
	}
}
