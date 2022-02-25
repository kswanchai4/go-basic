package main

import "fmt"

//scope
//globle varaible
var gVaraible int = 500

func main() {
	lVaraible := 40
	fmt.Println("Global =", gVaraible)
	fmt.Println("Local = ", lVaraible)
	anotherFunction()
}
func anotherFunction() {
	lVaraible := 20
	fmt.Println("Global =", gVaraible)
	fmt.Println("Local = ", lVaraible)
}
