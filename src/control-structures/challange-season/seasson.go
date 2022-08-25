package main

import "fmt"

func main() {
	var num int
	fmt.Println("Input a number to know the month: ")
	fmt.Scanln(&num)
	month, season := getMonth(num)
	fmt.Println("Month", "->", "Season")
	fmt.Println(month, "->", season)

}

func getMonth(month int) (string, string) {
	switch month {
	case 1:
		return "January", "Winter"
	case 2:
		return "February", "Winter"
	case 3:
		return "March", "Spring"
	case 4:
		return "April", "Spring"
	case 5:
		return "May", "Spring"
	case 6:
		return "June", "Summer"
	case 7:
		return "July", "Summer"
	case 8:
		return "August", "Summer"
	case 9:
		return "September", "Autumn"
	case 10:
		return "October", "Autumn"
	case 11:
		return "November", "Autumn"
	case 12:
		return "December", "Winter"
	default:
		return "Unknown", "Unknown"
	}
}
