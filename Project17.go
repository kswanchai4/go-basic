package main

import "fmt"

type Books struct {
	title    string
	author   string
	subtitle string
	price    float64
}

func main() {
	//Structer

	// var Golang Books

	// Golang.title = "Go Programming"
	// Golang.author = "LeoNadon"
	// Golang.subtitle = "LeoNadon Tutorial"
	// Golang.price = 199.99

	Golang := Books{
		title:    "Go Programming",
		author:   "LeoNadon",
		subtitle: "LeoNadon Tutorial",
		price:    199.99,
	}

	fmt.Println(Golang)
	fmt.Println(Golang.title)
}
