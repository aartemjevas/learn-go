package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Andreson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 21312,
		},
	}
	jim.updateName("Jimbo")
	jim.print()

}

func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
