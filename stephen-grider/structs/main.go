package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) update(newName string) {
	(*p).firstName = newName
}

func main() {
	neel := person{firstName: "Neel",
		lastName: "Shah",
		contactInfo: contactInfo{
			email: "neel@gmail",
			zip:   400071,
		},
	}
	neel.print()
	neel.update("Neel Kaushik")
	neel.print()
}
