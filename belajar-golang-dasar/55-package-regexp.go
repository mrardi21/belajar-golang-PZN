package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z]i)")

	fmt.Println(regex.MatchString("ardi"))
	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("aDi"))

	search := regex.FindAllString("ardi adi asi abi aDi", -1)
	fmt.Println(search)
}
