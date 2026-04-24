package main

import "fmt"

type MyInt int

func typeDefinition() {
	var a MyInt = 10
	var b int = 10

	// a = b ❌ Error (different types)}
	fmt.Println(a, b)
}

type myInt = int

func typeAlias() {
	var a myInt = 10
	var b int = 10

	fmt.Println(a, b)
	fmt.Println(a == b)
}

func structFn() {
	type Person struct {
		Name string
		Age  int
	}

	var arbaaz Person
	arbaaz.Name = "Arbaaz"
	arbaaz.Age = 22

	person := Person{
		Name: "meekail",
		Age:  22,
	}

	fmt.Println(arbaaz)
	fmt.Println(person)
}
