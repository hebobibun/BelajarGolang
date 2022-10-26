package main

import "fmt"

func main() {
	var sum = 11 + 10
	fmt.Println(sum)

	var a = 8
	var b = 4
	var c = a * b
	fmt.Println(c)

	var i = 5
	i *= 8  // i = i * 8

	fmt.Println(i)
}

// + - * / &

// Augmented Assignments
// a = a + 10    // a += 10
// a = a - 10    // a -= 10
// a = a * 10    // a *= 10
// a = a / 10    // a /= 10
// a = a % 10    // a %= 10

// Unary Operator
// ++   // a = a + 1
// --   // a = a - 1
// -    // Negative
// +    // Positive
// !    // Boolean Kebalikan