package main

import "fmt"

func main() {
	name := "Bob"

	switch name {
	case "Bob":
		fmt.Println("Hey Bob")
		fmt.Println("What's up")
	case "Budi":
		fmt.Println("Hey Budi")
	default:
		fmt.Println("Boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is too long!!")
	case false:
		fmt.Println("Correct!!")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Name is too long!!")
	case length > 5:
		fmt.Println("Name is a bit long")
	default:
		fmt.Println("Cool name!!")
	}
}
