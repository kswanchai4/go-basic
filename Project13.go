package main

import "fmt"

func main() {
	//x := make([]float64, 5)
	// slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// slice2 := append(slice1, 9, 10)
	// arr := [5]float64{1, 2, 3, 4, 5}
	// x := arr[0:5]
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice2)
}
