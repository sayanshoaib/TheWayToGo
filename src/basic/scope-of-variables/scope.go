package main

import "fmt"

var number int = 5 // number declared outside (global scope)

func main() {
	var decision bool = true // decision declared inside function(local scope)
	fmt.Println("Original Value of number: ", number)
	number = 10
	fmt.Println("New Value of number: ", number)
	fmt.Printf("Origanl Value of number: %d\n", number)

	fmt.Println("Value of decision: ", number)
	fmt.Println("Value of decision", decision)
	fmt.Println("Value from getNumber function: ", getNumber())

	fmt.Printf("New Value of number: %d\n", number)
	fmt.Printf("Value of decision: %t\n", decision)
}
func getNumber() int {
	return number
}
