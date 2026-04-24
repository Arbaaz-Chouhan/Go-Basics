package main

import (
	"fmt"
	"os"
)

func ReadFile() {
	data, err := os.ReadFile("test.tsx")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))

	dir, err := os.Getwd() // getwd check current directory

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Current directory:", dir)

	// check file info
	fileInfo, err := os.Stat("test.tsx")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File size:", fileInfo.Size())
	fmt.Println("File mode:", fileInfo.Mode())
	fmt.Println("File mod time:", fileInfo.ModTime())

	// create a new file
	file, _ := os.Create("test2.txt")

	fmt.Println(file.Name())
}
