package main

import "fmt"

func main() {
	type NoKTP string
	type married bool

	var noKTPhabib NoKTP = "9022893785"
	var marriedStatus married = false
	
	fmt.Println(noKTPhabib, marriedStatus)
}

// type declaration = membuat alias terhadap tipe data yang sudah ada