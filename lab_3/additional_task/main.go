package main

import (
	"fmt"
	"strconv"
	"log"
)

func isAllowedEquation(number int) bool {
	// return -2^15 <= number && number >= 2^15 - 1
	return number <= 2^15 - 1 && -2^15 <= number
}

func isPalindrome(text string) bool {

	intValue, err := strconv.Atoi(text)

	if err != nil {
		log.Fatalf("Error converting string to int: %v", err)
	}

	if !isAllowedEquation(intValue) {
		fmt.Println("\nBlocked by equation")
		return false
	}

	for k := 0; k <= int(len(text)/2); k++ {
		if text[k] == text[(len(text)-k)-1] {
			continue
		} else { return false }
	}

	// var 1
	// if text[0] == text[len(text) - 1] {
	// 	return true
	// }

	return true
}

func main() {
	// for {
	var input string
	
	fmt.Print("\nWrite your random number: ")
	fmt.Scan(&input)
	
	var result bool = isPalindrome(input)

	fmt.Printf("\nYour result: %t\n\n", result)
	// }
}