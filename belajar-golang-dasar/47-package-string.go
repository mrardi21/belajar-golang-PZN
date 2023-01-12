package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ardi Mardiansyah", "Ardi"))
	fmt.Println(strings.Contains("Ardi Mardiansyah", "Eko"))

	fmt.Println(strings.Split("Ardi Mardiansyah", " "))
	fmt.Println(strings.ToLower("Ardi Mardiansyah"))
	fmt.Println(strings.ToUpper("Ardi Mardiansyah"))
	fmt.Println(strings.ToTitle("Ardi Mardiansyah"))
	fmt.Println(strings.Trim("      Ardi Mardiansyah    ", " "))
	fmt.Println(strings.ReplaceAll("Ardi Mardiansyah", "Ardi", "Ardi Ganteng"))
}
