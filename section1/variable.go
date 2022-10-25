package main

import "fmt"

func main() {
	var name string

	name = "Habibullah"
	fmt.Println(name)

	name = "Muhammad Habibullah"
	fmt.Println(name)

	fmt.Println("")

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	fmt.Println("")

	school := "MIT"
	fmt.Println(school)

	school = "Harvard"
	fmt.Println(school)

	fmt.Println("")

	var (
		firstName = "Muhammad"
		lastName = "Habibullah"
	)

	fmt.Println(firstName, lastName)
}