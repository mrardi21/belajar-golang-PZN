package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHai(name string) {
	fmt.Println("Hello", name, ", My name is", customer.Name)
}

func main() {
	var ardi Customer

	ardi.Name = "Ardi"
	ardi.Address = "Bandung"
	ardi.Age = 20
	fmt.Println(ardi)

	budi := Customer{
		Name:    "Budi",
		Address: "Bandung",
		Age:     21,
	}
	fmt.Println(budi)

	joko := Customer{"Joko", "Jakarta", 21}
	fmt.Println(joko)

	joko.sayHai("Ardi")
}
