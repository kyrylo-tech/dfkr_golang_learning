package main

import "fmt"

func PrintSlice(prefix string, slice []int) {
	fmt.Printf("%s Slice: %v (len=%d) (cap=%d)\n", prefix, slice, len(slice), cap(slice))
}

func main()  {
	slice1 := []int{18}
	slice2 := []int{1, 28}
	slice3 := []int{76, 98, 58, 79, 68}

	PrintSlice("\nFirst", slice1)
	PrintSlice("Second", slice2)
	PrintSlice("Third", slice3)

	slice3 = slice3[:len(slice3)-2]

	PrintSlice("\nNew Third", slice3)

	allSlices := [][]int{slice1, slice2, slice3}

	fmt.Printf("\nAll Slices:\n")
	for i, slice := range allSlices {
		PrintSlice(fmt.Sprintf("#%d", i+1), slice)
	}

	fmt.Println()
}