package main

import "fmt"

func main() {
	var name1 = "Eko"
	var name2 = "Budi"

	var result = name1 == name2 // bool

	fmt.Println(result) // false

	var value1 = 100
	var value2 = 20

	fmt.Println(value1 > value2)  // true
	fmt.Println(value1 < value2)  // false
	fmt.Println(value1 == value2) // false
	fmt.Println(value1 != value2) // true
}
