package main

import "fmt"

const groupNumber int = 74
const variantNumber int = 1

func main() {
	name := "Kyrylo" 
	test := float64(variantNumber) * 3.14

	fmt.Printf("\nGroup number: %d\nVariant number: %d\nMy name: %s\nMy magic number: %.2f\n\n", groupNumber, variantNumber, name, test)
}