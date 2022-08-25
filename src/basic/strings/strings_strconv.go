package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "asasa"
	prefix := "a"
	// HasPrefix tests whether the string s begins with a prefix:
	fmt.Println(strings.HasPrefix(s, prefix))

	var str string = "This is an example of a string"
	fmt.Printf("T/F? \nDoes the string \"%s\" have prefix %s? ", str, "Th")

	fmt.Printf("Does the string \"%s\" have suffix %s? ", str, "ting")
	fmt.Printf("\n%t\n\n", strings.HasSuffix(str, "ting"))

	// Testing whether a string contains a substring
	var str2 string = "Hi, I'm mac, Hi."
	fmt.Println("The position of the first instance of\"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str2, "mac")) // Finding first occurrence
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi")) // Finding first occurrence
	fmt.Printf("The position of the last instance of \"Hi\" is:")
	fmt.Printf("%d\n", strings.LastIndex(str2, "Hi"))
	fmt.Printf("The position of the first instance of\"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger")) // Finding first occurrence

	var str3 string = "Hello, how is it going, Hugo?"
	var manyG = "ggggggg"
	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str3, "H"))
	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg")) // count occurences

	// Repeating a string#
	var origs string = "Hi there! "
	var newS string
	newS = strings.Repeat(origs, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)

	// Changing the case of a string
	var letter string = "aj amar mon valo nei"
	var up string = strings.ToUpper(letter)
	fmt.Printf("The original string is: %s\n", letter)
	fmt.Printf("The uppercas: %s", up)

	var orig string = "Hey, how are you George?"
	var lower string
	lower = strings.ToLower(orig) // changing to lower case
	fmt.Printf("The lowercase string is: %s\n", lower)

	var orig2 string = "666"
	var an int
	var new2 string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, _ = strconv.Atoi(orig2) // converting to number
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	new2 = strconv.Itoa(an) // converting to string
	fmt.Printf("The new string is: %s\n", new2)
}
