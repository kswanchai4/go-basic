package main

import "fmt"

func main() {
	summation(1, 2, 3, 4, 5)
}

func summation(args ...int) {
	var total int
	for _, n := range args {
		total += n
	}
	fmt.Println(total)
}
