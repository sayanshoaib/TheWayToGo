package main

import (
	"fmt"
)

//A label is a sequence of characters that identifies a location within a code.
//A code line which starts with a for,
//switch, or select statements can be preceded with a label of the form:

func main() {
Label1: // adding a label for location
	for i := 0; i <= 5; i++ { // outer loop
		for j := 0; j <= 5; j++ { // inner loop
			if j == 4 {
				break Label1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	fmt.Println()

Label2: // adding a label for location
	for i := 0; i <= 5; i++ { // outer loop
		for j := 0; j <= 5; j++ { // inner loop
			if j == 4 {
				continue Label2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	// goto IDENTIFIER

	i := 0
HERE:
	fmt.Printf("%d", i)
	i++
	if i == 5 {
		return
	}
	goto HERE

LABEL3: // adding a label for location
	for i := 0; i <= 5; i++ { // outer loop
		for j := 0; j <= 5; j++ { // inner loop
			if j == 4 {
				goto LABEL3 // jump to the label
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}
