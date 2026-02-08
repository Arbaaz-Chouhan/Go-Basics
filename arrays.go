package main

import "fmt"

func arraysFn() {
	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Pineapple"
	fruits[3] = "Mango"

	fmt.Println(fruits)    // [Apple Banana Pineapple Mango]
	fmt.Println(fruits[0]) /// Apple
	fmt.Println(fruits[1]) // Banana
	fmt.Println(fruits[2]) // Pineapple
	fmt.Println(fruits[3]) // Mango

	fmt.Println(len(fruits)) // 4
	fmt.Println(fruits[1:])  // [Banana Pineapple Mango]

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers) // [1 2 3 4 5]

	var vegetables = [4]string{"Potato", "Tomato", "Onion", "Cabbage"}
	fmt.Println(vegetables) // [Potato Tomato Onion Cabbage]

	marks := [3]int{80, 90, 70}

	for i := 0; i < 3; i++ {
		fmt.Println(marks[i]) // 80 90 70
	}

}

func arrayMap() {
	student := make(map[string]int)

	student["math"] = 90
	student["eng"] = 85

	details := map[int]string{
		1: "Arbaaz",
		2: "Ali",
	}

	fmt.Println(details)
	fmt.Println(details[1])
	fmt.Println(details[2])

	fmt.Println(student["math"])

	for k, v := range student {
		fmt.Println(k, v)
	}
}
