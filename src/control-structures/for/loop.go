package main

import (
	"fmt"
)

func main() {
	// for loop with 5 iterations
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}

	str := "Go is a beautiful language"
	fmt.Printf("The length of str is: %d\n", len(str))

	// for loop to find character at each position
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character at position %d is: %c \n", ix, str[ix])
	}

	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))

	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}

	var i int = 0
	for i < 5 {
		fmt.Printf("This is the %d iteration\n", i)
		i = i + 1
	}

	str3 := "Go is a beautiful programming language!"

	// for range
	for pos, char := range str3 {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str4 := "Chinese: 日本語"

	// for range
	for pos, char := range str4 {
		fmt.Printf("Character %c starts at byte position %d\n", char, pos)
	}
}
