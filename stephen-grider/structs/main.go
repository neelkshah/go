package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	neel := person{firstName: "Neel", lastName: "Shah"}
	fmt.Println(neel)

	var kaushik person
	kaushik.firstName = "Kaushik"
	kaushik.lastName = "Shah"
	fmt.Printf("%+v", kaushik)
}
