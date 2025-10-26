package main

import "fmt"

func filterInts(nums []int, pred func(int) bool) []int {
	result := make([]int, 0)
	for _, n := range nums {
		if pred(n) {
			result = append(result, n)
		}
	}
	return result
}

func main() {
	variant := 1
	nums := []int{6, 8, 14, 11, 32, 23, 21, 29, 50, 9}

	fmt.Println("Початковий зріз:", nums)

	evens := filterInts(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Парні числа:", evens)

	greater := filterInts(nums, func(n int) bool { return n > variant })
	fmt.Println("Більші за варіант:", greater)
}
