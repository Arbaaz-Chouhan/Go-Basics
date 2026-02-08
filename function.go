package main

import "fmt"

func function() {
	fmt.Println("First Function")
}

func forloopFunction() {
	for i := 0; i < 10; i++ {
		// fmt.Println("Hello World")
		fmt.Println(i)
	}
}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a, b int) (int, int) {
	return a / b, a % b
}

// Function without pointer
func functionWithoutPointer(a int) {
	a = 20
}

// Function with pointer
func functionWithPointer(a *int) {
	*a = 20
}
