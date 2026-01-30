package main

import (
	"fmt"
	"strconv"
	"math"
) 

func main() {
	var groupNumber int8 = 4
	
	var variantType func(float64) float64 = math.Cos 
	var variantName string = "cosinus"
	
	if groupNumber % 2 == 0 {
		variantType = math.Sin
		variantName = "sinus"
	}

	var result float64 = variantType(float64(groupNumber))

	fmt.Printf("\nMy group list: %s\n", strconv.FormatInt(int64(groupNumber), 10))
	fmt.Println("My function:", variantName)
	fmt.Printf("My result: %s\n\n", strconv.FormatFloat(result, 'f', 2, 64))
}