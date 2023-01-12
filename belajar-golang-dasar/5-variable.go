package main

import "fmt"

func main() {
	var name string

	name = "ardi mardiansyah"
	fmt.Println(name)

	name = "sakinah mw"
	fmt.Println(name)

	var alamat = "Bandung, Jawa Barat"
	fmt.Println(alamat)

	var age int8 = 25
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	// Delarasi multiple variable
	var (
		firstName = "Ardi"
		Lastname  = "Mardiansyah"
	)

	fmt.Println(firstName, Lastname)
}
