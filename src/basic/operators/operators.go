package main

import "fmt"

func main() {
	var a float32 = 1 / 4
	b := 1
	a = float32(b)
	fmt.Println(a)

	b3 := 10 > 5 // greater than operator
	fmt.Println(b3)
	b3 = 10 < 5
	fmt.Println(b3) // less than operator
	fmt.Println(b3)
	b3 = 5 <= 10 // less than equal to
	fmt.Println(b3)
	b3 = 10 != 10 // not equal to
	fmt.Println(b3)

	c3 := 10 > 5 && 7 < 15 // AND operator
	fmt.Println(c3)
	c3 = 10 < 5 || 2 > 7 // OR operator
	fmt.Println(c3)
	c3 = !c3 // NOT operator
	fmt.Println(c3)
}
