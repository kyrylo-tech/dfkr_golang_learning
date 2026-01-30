package main

import "fmt"

// Group = 79 | Variant = 316 | Course = -0.44

func main() {
	myArray := [5]float64{1.0, 2.0, 79.0, 316.0, -.044}
	variant := 1.0

	fmt.Printf("\nVariant: %.f", variant)

	fmt.Printf("\n\nOld array: %v", myArray)

	myArray[0] = 74.0
	myArray[1] = 4.0

	for i := 0; i < len(myArray); i++ {
		myArray[i] += variant
	}

	fmt.Printf("\nNew array: %v\n\n", myArray)
}