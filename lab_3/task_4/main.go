package main

import "fmt"

func main() {
	defer fmt.Println("Task made by Kyrylo from 74 group\n")

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

	for course, groups := range matrix {
		for _, group := range groups {
			if group == groupNumber {
				targetCourse = course
			}
		}
	}

	if targetCourse == 0 {
		fmt.Println("\nNo one course not found with this group name\n")
	} else {
		fmt.Printf("\nCourse of group number '%d' = %d\n\n", groupNumber, targetCourse)
	}
}