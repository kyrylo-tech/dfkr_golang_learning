package main

import "fmt"

type ContactInfo struct {
	Email string
	Phone string
}

type Address struct {
	City string
	Street string
}

type Employee struct {
	FirstName string
	LastName string
	Position string
	Contact ContactInfo
	Address
}

func PrintEmployeeInfo(empl Employee) {
	fmt.Println("\n--- Employee Information ---")
	fmt.Printf("Fullname: %s %s", empl.FirstName, empl.LastName)
	fmt.Printf("\nPosition: %s", empl.Position)
	fmt.Printf("\nEmail: %s\nCity: %s\n\n", empl.Contact.Email, empl.Address.City)
}

func main() {
	newEmployee := Employee{
		FirstName: "Danylo",
		LastName: "Kovalenko",
		Position: "Programmer",
		Contact: ContactInfo{
			Email: "kovalenko@ukr.net",
			Phone: "+380000000001",
		},
		Address: Address{
			City: "Dnipro",
			Street: "st. Stepana Bandery",
		},
	}
	PrintEmployeeInfo(newEmployee)
}