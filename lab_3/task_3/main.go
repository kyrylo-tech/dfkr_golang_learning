package main

import "fmt"

func main() {
	matrix := map[int][]int{
		1: {87, 88, 89, 90, 91},
		2: {81, 82, 83, 84, 85, 86},
		3: {75, 76, 77, 78, 79, 80},
		4: {69, 70, 71, 72, 73, 74},
	}

	var groupNumber int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&groupNumber)

	var targetCourse int

search:
	for course, groups := range matrix {
		for _, group := range groups {
			if group == groupNumber {
				switch course {
				case 1:
					targetCourse = 1
				case 2:
					targetCourse = 2
				case 3:
					targetCourse = 3
				case 4:
					targetCourse = 4
				}
				break search
			}
		}
	}

	fmt.Printf("\nCourse of group number '%d' = %d\n\n", groupNumber, targetCourse)
}
