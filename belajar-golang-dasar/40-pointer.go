package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// pass by value
	// address2 := address1
	// pass by reference
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"
	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)
}
