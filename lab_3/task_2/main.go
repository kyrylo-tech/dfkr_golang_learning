package main

import (
	"fmt"
	"math"
	"strconv"
)

func sqrt(num int, lenght int) string {
	var result float64 = math.Sqrt(float64(num))
	var newLen float64 = float64(lenght)

	switch {
	case result == newLen:
		return strconv.FormatFloat(result, 'f', 2, 64) + " & " + strconv.FormatFloat(newLen, 'f', 2, 64)
	case result > newLen:
		return strconv.FormatFloat(result + newLen, 'f', 2, 64) + "s"
	case result < newLen:
		return strconv.FormatFloat(result * newLen, 'f', 2, 64) + "n"
	default:
		return "none"
	}
}

func main() {
	var groupNumber int = 74
	var myNameLenght int = 6

	var myResult string = sqrt(groupNumber, myNameLenght)

	fmt.Printf("\nMy sqrt of group number %d is %s\n\n", groupNumber, myResult)

}