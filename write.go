package main

import (
	"fmt"
	"os"
)

func WriteFile() {
	data := []byte("Hello, this is written from Go!")

	err := os.WriteFile("test.txt", data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File written successfully")
}

func CreateFile() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	file.WriteString("Hello from os.Create()\n")
	file.WriteString("Second line\n")

	fmt.Println("File created and written")
}
