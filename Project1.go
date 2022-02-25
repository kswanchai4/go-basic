package main

import "fmt"

// 	ตัวแปร

/*
	การประกาศตัวแปร
*/

func main() {
	var myAge int = 38
	myAge2 := 40.1
	var something bool = true
	text := "Leo Nadon Story"
	age1, age2 := 50, 82

	fmt.Println("Value Variable = ", myAge)
	fmt.Println("Value = ", myAge2)
	fmt.Println("Value = ", something)
	fmt.Println(text)
	fmt.Println(age1, age2)
}
