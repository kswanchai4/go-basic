package main

import "fmt"

func main() {
	dosumething("Leo")
	dosumething("Nadon")
	addition(5, 10)
	result := addition2(20, 30)
	fmt.Println(result * 10)
}
func dosome() {
	fmt.Println("OK")
}
func dosumething(str string) {
	fmt.Println(str)
}
func addition(a int, b int) {
	fmt.Println(a + b)
}
func addition2(a int, b int) int {
	output := a + b
	return output
}
