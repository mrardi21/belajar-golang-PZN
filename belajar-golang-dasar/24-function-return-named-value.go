package main

import "fmt"

func getFullName() (firstName, lastName string) {
	firstName = "Ardi"
	lastName = "Mardiansyah"
	return
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
