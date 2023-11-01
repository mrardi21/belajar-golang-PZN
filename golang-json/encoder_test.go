package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestSreamEncoder(t *testing.T) {
	writer, _ := os.Create("Customer_output.json")
	encoder := json.NewEncoder(writer)

	customer := &Customer{
		Firstname: "Ardi",
		LastName:  "Mardiansyah",
	}
	encoder.Encode(customer)

	fmt.Println(customer)
}
