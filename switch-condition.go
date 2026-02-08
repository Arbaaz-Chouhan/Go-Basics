package main

import "fmt"

func switchCondition() {
	// example 1
	// age := 18

	// switch age {
	// case 18:
	// 	fmt.Println("You can a vote ")
	// default:
	// 	fmt.Println("You can't vote")

	// }

	// example 2
	// day := 3

	// switch day {
	// case 1:
	// 	fmt.Println("Monday")
	// case 2:
	// 	fmt.Println("Tuesday")
	// case 3:
	// 	fmt.Println("Wednesday")
	// default:
	// 	fmt.Println("Invalid day")
	// }

	// example 3
	lang := "go"

	switch lang {
	case "go":
		fmt.Println("Golang")
	case "js":
		fmt.Println("JavaScript")
	default:
		fmt.Println("Unknown language")
	}

}
