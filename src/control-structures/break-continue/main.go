package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			fmt.Printf("%d", j)
		}
		fmt.Println(" ", " -> ", i)
	}

	for i := 0; i < 5; i++ {
		if i > 3 {
			fmt.Println(i)
			break
		}
		fmt.Println("*")
	}

	for i := 0; i < 10; i++ { // for loop
		if i == 5 {
			continue // continuing the loop
		}
		fmt.Printf("%d", i)
		fmt.Print(" ")
	}
}
