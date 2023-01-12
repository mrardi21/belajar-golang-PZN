package main

import (
	"fmt"
)

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}

	fmt.Println("Run application")
}

func main() {

	runApp(true)
	fmt.Println("test run")
}
