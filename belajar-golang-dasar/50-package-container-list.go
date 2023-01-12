package main

import (
	"container/list"
	"fmt"
)

type Customer struct {
	Name string
}

func main() {
	data := list.New()
	data.PushBack("Ardi")
	data.PushBack("Mardiansyah")
	data.PushBack("Ganteng")
	data.PushFront("Hi")
	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)
	// fmt.Println(data.Len())

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	// for element := data.Back(); element != nil; element = element.Prev() {
	// 	fmt.Println(element.Value)
	// }

	data2 := list.New()
	data2.PushBack(Customer{Name: "Ardi"})

	for element := data2.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
