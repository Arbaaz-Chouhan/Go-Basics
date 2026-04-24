// ek struct ko dusre struct me directly include karke uski properties aur methods use karna

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Admin struct {
	User
	Role string
}

func embedding() {

	admin := Admin{
		User: User{
			Name: "Arbaaz",
			Age:  22,
		},
		Role: "Developer",
	}

	fmt.Println(admin.Name)
	fmt.Println(admin.Age)
	fmt.Println(admin.Role)
}

type User2 struct {
	Name string
}

func (u User) greet() {
	fmt.Println("Hello", u.Name)
}

type Admin2 struct {
	User
}

func embeddingMethod() {
	a := Admin{
		User: User{Name: "Arbaaz"},
	}

	a.greet() // method inherited
}
