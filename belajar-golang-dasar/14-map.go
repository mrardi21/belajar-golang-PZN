package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "ardi",
		"address": "Bandung",
	}
	person["job"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["job"])

	// var books map[string]string = make(map[string]string)
	books := make(map[string]string)
	books["title"] = "Belajar Golang"
	books["author"] = "Ardi"
	books["ups"] = "Salah"

	fmt.Println(books)
	delete(books, "ups")
	fmt.Println(books)
}
