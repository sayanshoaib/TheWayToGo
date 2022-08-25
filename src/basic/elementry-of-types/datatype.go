package main

import (
	"fmt"
	"math/rand"
)

// elementary data types or primate data types
/*
The three main elementary types in Go are:
-Boolean
-Numeric
-Character

*/

func main() {
	var money uint = 89
	fmt.Println(money)

	var c1 complex64 = 5 + 10i // Declaring complex num (real +imaginary(¡))
	fmt.Printf("The value is: %v\n", c1)

	a := rand.Int()   // generates a random number
	b := rand.Intn(8) // generates a random number in [0, n)
	fmt.Printf("a is: %d\n", a)
	fmt.Printf("b is: %d\n", b)

	var ch1 int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch1, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch1, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch1, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch1, ch2, ch3)   // UTF-8 code point

}
