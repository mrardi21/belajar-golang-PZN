package main

import (
	"fmt"
)

func main() {
	name := "Ardi"

	switch name {
	case "Ardi":
		fmt.Println("Hello Ardi")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, kenalan donk")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Benar")
	case false:
		fmt.Println("Salah")
	}
}
