package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", hostname)
	} else {
		fmt.Println("Error", err)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println(username)
	fmt.Println(password)

	file, err := os.Open("test.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	} else {
		// fmt.Println(file)
		data := make([]byte, 100)
		count, err := file.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	}

}
