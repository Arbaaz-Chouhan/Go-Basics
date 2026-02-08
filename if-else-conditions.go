package main

import "fmt"

func ifElseConditions() {
	number := 10

	if number > 0 {
		fmt.Println("Positive")
	} else {
		fmt.Println("Negative")
	}
}

func equals() {
	age := 18

	if age == 18 {
		fmt.Println("You can a vote ")
	} else {
		fmt.Println("You can't vote")
	}
}

func marks() {
	marks := 80

	if marks >= 80 {
		fmt.Println("Grade A")
	} else if marks >= 60 {
		fmt.Println("Grade B")
	} else {
		fmt.Println("Grade C")
	}
}
