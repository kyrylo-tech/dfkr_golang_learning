package main

import "fmt"

const nameCharacters int = 6
const groupNumber int = 1

func main() {
	var tempSum int = nameCharacters + groupNumber
	var realSum int = 0

	for i := 0; i < tempSum; i++ {
		realSum += i
	}

	fmt.Printf("\nName characters amount: %d\nGroup journal number: %d\nNumbers sum from 1 to %d: %d\n\n", nameCharacters, groupNumber, tempSum, realSum)
}