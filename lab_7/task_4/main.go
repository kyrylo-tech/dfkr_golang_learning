package main

import (
	"fmt"
	"time"
)

func sendAfter(delay time.Duration, ch chan string, msg string) {
	time.Sleep(delay)
	ch <- msg
}

func main() {
	ch := make(chan string)

	go sendAfter(2*time.Second, ch, "Дані з каналу після 2 секунд")
	go sendAfter(4*time.Second, ch, "Дані з каналу після 4 секунд")

	select {
	case msg := <-ch:
		fmt.Println("Отримано:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Таймаут: відповідь не надійшла протягом 3 секунд")
	}
}
