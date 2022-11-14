package main

import "fmt"

func main() {
	months := [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	slice1 := months[5:9] //[low:high]

	fmt.Println(months)
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // length of slice
	fmt.Println(cap(slice1)) // capacity (from low to end of arrays)

	months[2] = "change array"
	fmt.Println(months)

	slice1[2] = "change slice" // low is index 0 of slice1, so [2] in slice1 is [7] in months
	fmt.Println(slice1)
	fmt.Println(months)

	//////////////////////////////////////////

	// Playing with append()

	days := [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	fmt.Println("")
	fmt.Println("days      =", days)
	fmt.Println("")

	daySlice1 := days[4:]
	daySlice1[0] = "NewFriday"
	daySlice1[1] = "NewSaturday"

	fmt.Println("daySlice1 =", daySlice1)
	fmt.Println("days      =", days)
	fmt.Println("")

	daySlice2 := append(daySlice1, "Minggu")
	fmt.Println("daySlice2 =", daySlice2)
	fmt.Println("")

	fmt.Println("daySlice1 =", daySlice1)
	fmt.Println("")

	fmt.Println("days      =", days)
	fmt.Println("")

	//////////////////////////////////////////

	// Playing with make()

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Jakarta"
	newSlice[1] = "Bandung"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	newAppend := append(newSlice, "Semarang", "Jogja", "Surabaya")

	fmt.Println(newAppend)

	fmt.Println("")

	//////////////////////////////////////////

	// Playing with copy()

	copySlice := make([]string, len(newAppend), cap(newAppend)) // make sure it has same len & cap
	copySlice2 := make([]string, 3, cap(newAppend))             // if so, it will return only 3 of newAppend data

	copy(copySlice, newAppend)  // (destination, source)
	copy(copySlice2, newAppend) // (destination, source)

	fmt.Println(copySlice)
	fmt.Println(copySlice2)

	fmt.Println("")

	//////////////////////////////////////////

	// BE CAREFUL WHEN YOU DECLARE ARRAY

	anArray := [...]string{"oke", "sip", "mantap"} // it's an array
	anArray2 := [3]string{"oke", "sip", "mantap"}  // it's an array
	aSlice := []string{"oke", "sip", "mantap"}     // it's a slice

	fmt.Println(anArray)
	fmt.Println(anArray2)
	fmt.Println(aSlice)

}
