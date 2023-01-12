package main

import "fmt"

func main() {
	var name [2]string

	name[0] = "Ardi"
	name[1] = "Mardiansyah"

	fmt.Println(name[0])
	fmt.Println(name[1])

	var values = [3]int{
		90, 88, 95,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
