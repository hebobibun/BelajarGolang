package main

import "fmt"

func main() {
	grade := 85

	if grade >= 80 {
		fmt.Println("Congratulations!!")
	} else if grade >= 60 && grade < 80 {
		fmt.Println("keep it up!!")
	} else {
		fmt.Println("Sorry your grade is too low")
	}

	name := "bob"

	if name == "Habib" {
		fmt.Println("Hey", name, "!!")
	} else if name == "bob" {
		fmt.Println("Hey", name, "!!")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Your name is correct!!")
	}
}
