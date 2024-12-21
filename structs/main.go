package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	age int
}

type personExtended struct {
	firstName string
	lastName string
	age int
	contact contactInfo
}

func main() {
	p1 := person{} // Creates as a zero-value ("", 0 or false)  
	p2 := person{"Alex", "Anderson", 27}
	p3 := person{firstName: "Tim", age: 68, lastName: "Cook"}

	fmt.Println(p1, p2, p3)
	fmt.Printf("%+v", p2) // To print with their values name

	p1.firstName = "Joaquin" // Easily access and modify a struct field
	p1.lastName = "Jimenez"
	p1.age = 27

	fmt.Println(p1)

	jim := personExtended{
		firstName: "Jim",
		lastName: "Halpert",
		age: 45,
		contact: contactInfo{
			email: "jimothy@scranton.com",
			zipCode: 34000,
		},
	}

	jim.print()

	p1.firstName = "Emanuelita" // We can override the value of first name
	p1.print()

	p2.updateName("Albert") // This won't change the name since we are passing a copy
	p2.print() // Still old name

	p2Pointer := &p2 // With & we are asking for the access in the memory where the var is at
	// p2Pointer is not an object, it's a memory address, a pointer
	p2Pointer.updateNameExtendedPointer("Albert")
	p2.print()

	// Trying to change the name again using the shortcut and passing directly p2
	// we can send directly a variable to a pointer and it will do the conversion itself
	p2.updateNameIdealfunction("Another")
	p2.print()
}

func (p personExtended) print() {
	fmt.Println(p)
}

func (p person) print() {
	fmt.Println(p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// With * we are accessing to the value of a memory address (a pointer)
func (pp *person) updateNameExtendedPointer(newFirstName string) {
	(*pp).firstName = newFirstName
}

// Shortcut to manipulate structs without converting to pointers and values
func (pp *person) updateNameIdealfunction(newFirstName string) {
	pp.firstName = newFirstName
}