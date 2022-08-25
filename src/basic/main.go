package main // making package for standalone executable
import "fmt" // importing a package

func main() { // making an entry point
	// printing using fmt functionality
	fmt.Println("Hello World")
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")

	// Type casting
	var number float32 = 5.2
	fmt.Println(number)
	fmt.Println(int(number))

	// Constants
	// Implicit Typing
	const PI = 3.14159
	const B = "hello"

	// Explicit typing
	const URL string = "https://aplhaGeeks.com"

	// Typed and untyped constants
	var n int
	fmt.Println(n)

	//const (
	//	UNKNOWN = 0
	//	FEMALE  = 1
	//	MALE    = 2
	//)
	type Gender int
	const (
		UNKNOWN = iota
		FEMALE  = iota
		MALE    = iota
	)

	fmt.Println(Gender(FEMALE))
} // exiting the program
