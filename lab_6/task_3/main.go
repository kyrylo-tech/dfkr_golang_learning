package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return fmt.Sprintf("Привіт, мене звуть %s.", p.Name)
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) GetEmployeeID() string {
	return fmt.Sprintf("ID співробітника: %s", e.EmployeeID)
}

func (e Employee) Greet() string {
	return fmt.Sprintf("Привіт, я співробітник %s з ID %s.", e.Name, e.EmployeeID)
}

func main() {
	emp := Employee{
		Person:     Person{Name: "Кирило", Age: 18},
		EmployeeID: "E-123",
	}

	fmt.Println("Ім’я:", emp.Name)
	fmt.Println("Вік:", emp.Age)
	fmt.Println(emp.GetEmployeeID())

	fmt.Println("\nGreet (Employee):", emp.Greet())
	fmt.Println("Greet (Person):", emp.Person.Greet())
}
