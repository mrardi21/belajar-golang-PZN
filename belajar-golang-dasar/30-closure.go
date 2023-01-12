package main

import "fmt"

func main() {
	counter := 0
	name := "Ardi"

	increment := func() {
		fmt.Println("increment")
		counter++
		name := "Mardi"
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
