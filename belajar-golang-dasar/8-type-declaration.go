package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpArdi NoKTP = "3411161020"
	var marriedStatus Married = true

	fmt.Println(noKtpArdi)
	fmt.Println(marriedStatus)
}
