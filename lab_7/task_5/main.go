package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range jobs {
		square := num * num
		fmt.Printf("Worker #%d: %d^2 = %d\n", id, num, square)
		results <- square
	}
}

func main() {
	numbers := []int{6, 8, 14, 11, 32, 23, 21, 29, 50, 9}

	jobs := make(chan int, len(numbers))
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for _, num := range numbers {
		jobs <- num
	}
	close(jobs)

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println("Результат:", r)
	}
}
