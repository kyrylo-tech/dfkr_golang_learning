package main

import "fmt"

func createCounter(start int) func() int {
	curr := start
	return func() int {
		val := curr
		curr++
		return val
	}
}

func main() {
	variant := 1
	counterFromZero := createCounter(0)
	counterFromVar := createCounter(variant)

	fmt.Println("Лічильник від 0:")
	for i := 0; i < 4; i++ {
		fmt.Println(counterFromZero())
	}

	fmt.Println("\nЛічильник від варіанта:")
	for i := 0; i < 3; i++ {
		fmt.Println(counterFromVar())
	}

	fmt.Println("\nЩе один виклик першого:", counterFromZero())
}
