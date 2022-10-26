package main

import "fmt"

func main() {
	const firstName = "Muhammad "
	const lastname = "Habibullah"
	const value = 250

	fmt.Println(firstName)
	fmt.Println(value)

	const (
		country = "Indonesia"
		continent = "ASIA"
	)

	fmt.Println("")
	fmt.Println(country)
	fmt.Println(continent)

}

// Constant boleh tidak digunakan secara langsung, golang won't be complaining