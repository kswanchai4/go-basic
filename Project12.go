package main

import "fmt"

func main() {

	// var x [5]int
	// x[0] = 10
	// x[1] = 20
	// x[2] = 30
	// x[3] = 40
	// x[4] = 50
	x := [5]float64{10, 20, 30, 40, 50}
	var total float64
	for _, value := range x {
		total += value
	}

	fmt.Println(total)
	fmt.Println(total / float64(len(x)))
}
