package main

import "fmt"

func main() {
	var firstName string = "Shoaib"
	fmt.Println(firstName)
	var lastName string = "Hasan"
	var fullName = firstName + " " + lastName
	fmt.Println(fullName)
	var rawString string = `This is a raw string \n`
	fmt.Println(rawString)

	s := "hel" + "lo "
	s += "World"
	fmt.Println(s)

}
