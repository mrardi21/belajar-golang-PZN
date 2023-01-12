package main

import "fmt"

func getFullName() (string, string, string) {
	return "Mr.", "Ardi ", "Mardiansyah"
}

func main() {
	_, firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
