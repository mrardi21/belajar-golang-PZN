package main

import "fmt"

type Filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ardi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	// fmt.Println(sayGoodBye("Ardi"))
}
