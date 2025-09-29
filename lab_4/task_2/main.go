package main

import ("fmt")

type Person struct {
	FirstName string
	LastName string
	GroupNumber int
	Variant int
}

func modify(person *Person) {
	var sum int = (person.GroupNumber) + (person.Variant)
	person.Variant = sum
}

func main() {
	me := Person{
		FirstName: "Kyrylo",
		LastName: "Bitsay",
		GroupNumber: 74,
		Variant: 1,
	}

	fmt.Println(me)
	modify(&me)
	fmt.Println(me)
}