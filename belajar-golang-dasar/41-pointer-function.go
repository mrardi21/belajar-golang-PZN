package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	alamat := Address{
		City:     "Bandung",
		Province: "Jawa Barat",
		Country:  "",
	}
	// Pass By Value
	// ChangeCountryToIndonesia(alamat)
	// Pass By Reference
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
