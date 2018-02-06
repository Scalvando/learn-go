package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	/*
		//alex := person{firstName: "Alex", lastName: "Anderson"}
		var alex person

		alex.firstName = "Alex"
		alex.lastName = "Anderson"

		fmt.Println(alex)
		fmt.Printf("%+v\n", alex)
	*/

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")

	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
