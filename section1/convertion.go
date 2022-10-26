package main

import "fmt"

func main() {
	var nilai32 int32 = 130
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // Bila telah mencapai nilai Max, maka akan langsung balik ke nilai Min

	fmt.Println("")
	
	var name = "Habib"
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}

