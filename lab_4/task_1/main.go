package main

import (
	"fmt"
	"math"
)

func Calculations(val1 *int, val2 *int, val3 *float64) {
	var group int = *val1

	*val1 = *val1 + (*val2) + int(*val3)
	*val2 = *val1 * (*val2) * int(*val3)

	if group % 2 == 0 {
		*val3 = math.Sin(float64(*val1))
	} else {
		*val3 = math.Cos(float64(*val2))
	}
}

func main() {
	var group int = 74
	var variant int = 1
	var course float64 = 4.0

	fmt.Printf("\nOld:\nGroup = %d | Variant = %d | Course = %.2f", group, variant, course)

	Calculations(&group, &variant, &course)

	fmt.Printf("\n\nNew:\nGroup = %d | Variant = %d | Course = %.2f\n\n", group, variant, course)
}