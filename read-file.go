package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	// data, err := os.ReadFile("test.tsx")

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println(string(data))

	dir, err := os.Getwd() // getwd check current directory

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Current directory:", dir)
}
