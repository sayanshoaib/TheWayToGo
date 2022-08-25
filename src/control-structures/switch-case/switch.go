package main

import (
	"fmt"
	"strings"
)

func main() {
	var num1 int = 100
	// Adding switch on num1
	switch num1 {
	case 98, 99:
		fmt.Println("It's is equal to 98 or 99")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not between 98 to 100")

	}

	var gender string

	fmt.Println("Input your gender: ")
	fmt.Scan(&gender)
	gender = strings.ToLower(gender)
	switch gender {
	case "male":
		fmt.Println("Validating male candidate")
	case "female":
		fmt.Println("Validating female candidate")
	default:
		fmt.Println("can't validate")
	}

	var num2 int = 100
	switch {
	case num2 < 100:
		fmt.Println("num2 is less than 100")
	case num2 == 100:
		fmt.Println("num2 is equal to 100")
	case num2 > 100:
		fmt.Println("num2 is greater than 100")
	default:
		fmt.Println("Janina")

	}

	switch num1 := 100; {
	case num1 < 0:
		fmt.Println("Number is negative")

	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")

	default:
		fmt.Println("Number is 10 or greater")
	}

}
