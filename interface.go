// 🔵 Interface kya hota hai?

// 👉 Interface ek contract hota hai.
// Matlab:

// “Jo bhi type ye methods implement karega, wo is interface ka part ban jayega.”

// Go me interface implicitly implement hota hai
// (jaise OOP languages me implements likhna padta hai, Go me nahi ❌)

package main

import "fmt"

// Interface
type Animal interface {
	Speak() string
}

// Struct 1
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

// Struct 2
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

func interfacefn() {
	var a Animal

	a = Dog{}
	fmt.Println(a.Speak())

	a = Cat{}
	fmt.Println(a.Speak())
}
