package main

import "fmt"

func main() {
	myArray := [10]int{6, 8, 14, 11, 32, 23, 21, 29, 50, 9}

	sl_a := myArray[:5]
	sl_b := myArray[3:7]
	sl_c := myArray[8:]

	fmt.Printf("\nStart array: %v\n\nSlice a: %v\nSlice b: %v\nSlice c: %v\n", 
		myArray, sl_a, sl_b, sl_c)

	sl_a[1] = 11
	sl_b[0] = 22
	sl_c[len(sl_c)-1] = 33

	newArray := append(sl_a, sl_b...)
	newArray = append(newArray, sl_c...)

	fmt.Printf("\nFinal array: %v\n\n", newArray)
}