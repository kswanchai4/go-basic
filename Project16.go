package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("value is %d\n", x)
	fmt.Printf("Address x variable %x\n", &x)
	var p *int
	p = &x // XXX:
	fmt.Printf("Pointer P = %x\n", p)
	*p = 20
	fmt.Printf("value x variable %d\n", x)
}
