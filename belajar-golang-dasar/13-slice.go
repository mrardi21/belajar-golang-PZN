package main

import "fmt"

func main() {
	month := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))

	var slice2 = month[2:4]
	fmt.Println(slice2)
	fmt.Println(cap(slice2))

	var slice3 = append(slice2, "Ardi")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(month)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Ardi"
	newSlice[1] = "Mardiansyah"

	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
