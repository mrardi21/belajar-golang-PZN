package main

import (
	"fmt"
)

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	ardi := Man{"Ardi"}
	ardi.Married()
	fmt.Println(ardi.Name)
}
