package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	derek := person{
		firstName: "Derek",
		lastName: "Garnett",
		contactInfo: contactInfo{
			email: "drkgrntt@gmail.com",
			zipCode: 64055,
		},
	}

	derek.updateFirstName("Eric")
	derek.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}
