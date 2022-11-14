package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Habib"
	names[1] = "Habibullah"
	names[2] = "Hebobibun"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		15,
		25,
		35,
	}

	values[2] = 10 // change value

	fmt.Println(values)
	fmt.Println(values[0]) // show array index 0
	fmt.Println(values[1]) // show array index 1
	fmt.Println(values[2]) // show array index 2

	fmt.Println(len(names))              // show the length of array
	fmt.Println(len(values))             // show the length of array
	fmt.Println(len(names), len(values)) // show the length of array

	var test [10]string

	fmt.Println(len(test)) // works fine without values/data
}
