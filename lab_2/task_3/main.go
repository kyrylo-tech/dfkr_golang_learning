package main

import "fmt"

func birthday_math(dd int64, mm int64, yyyy int64) (int64, int64) {
	if (dd <= 0 || dd > 31) { return 0, 0 }
	if (mm <= 0 || mm > 12) { return 0, 0 }

	var sum int64 = dd + mm + yyyy
	var dob int64 = dd * mm * yyyy

	return sum, dob
}

func main() {
	var sum, dob int64 = birthday_math(11, 11, 1111)

	fmt.Printf("\nSumma: %d\nDobutok: %d\n\n", sum, dob)
}