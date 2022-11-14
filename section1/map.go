package main

import "fmt"

func main() {
	person := map[string]string{ // map[Key]Value || map[KeyDataType]ValueDataType
		"Name":    "Habibullah",
		"Country": "Indonesia",
	}

	person["Country"] = "Wakanda"  // CHANGE
	person["Title"] = "Programmer" // ADD NEW

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["Name"])
	fmt.Println(person["Country"])
	fmt.Println(person["Title"])

	book := make(map[string]string)
	book["title"] = "Golang Book"
	book["author"] = "Habibullah"
	book["oops"] = "wrong"

	fmt.Println(book)
	fmt.Println(len(book)) // 3

	delete(book, "oops") // delete(map, key)

	book["pages"] = "120"
	book["color"] = "blue"

	fmt.Println(book)
	fmt.Println(len(book)) // 4
}
