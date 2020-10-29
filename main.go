package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func main() {
	alex := person{
		firstName:   "Alex",
		lastName:    "Anderson",
		contactInfo: contactInfo{"abc@gmail.com", 555},
	}
	alex.print()

	alex.updateName("Alexxx")
	alex.print()
	unknown := person{}
	unknown.print()

	InterfacesMain()
	HTTPMain()
	ChannelsMain()

	fmt.Println("end")
}
