package main

import "fmt"

type Info struct {
	FirstName string
	LastName string
	Group int
	Variant int
}

type Student struct {
	Info Info
	Value int
}

func main()  {
	myArray := [10]int{6, 8, 14, 11, 32, 23, 21, 29, 50, 9}
	thisSlice := myArray[:4]

	students := [4]Student{
		Student{
			Info: Info{
				FirstName: "Petro",
				LastName: "Petrenko",
				Group: 75,
				Variant: 2,
			},
			Value: thisSlice[0],
		},
		Student{
			Info: Info{
				FirstName: "Petro",
				LastName: "Petrenko",
				Group: 1,
				Variant: 2,
			},
			Value: thisSlice[1],
		},
		Student{
			Info: Info{
				FirstName: "Petro",
				LastName: "Petrenko",
				Group: 1,
				Variant: 2,
			},
			Value: thisSlice[2],
		},
		Student{
			Info: Info{
				FirstName: "Petro",
				LastName: "Petrenko",
				Group: 1,
				Variant: 2,
			},
			Value: thisSlice[3],
		},
	}

	for i:=0; i<len(students); i++ {
		s := students[i]
		fmt.Printf("Element %d:\n  Information: %s %s, group %d, variant %d\n  Variant: %d\n", i+1, s.Info.FirstName, s.Info.LastName, s.Info.Group, s.Info.Variant, s.Value)
	}
}