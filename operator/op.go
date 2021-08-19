package main

import "fmt"

func main() {
	a := 4
	b := 2
	// var c intc = 5
	// var d = 6
	// var e = 3.14
	// var f int = 8
	fmt.Printf("%v\n", a&b)
	fmt.Printf("%v\n", a|b)
	fmt.Println("a&b=", a&b)
	fmt.Println("a|b=", a|b)
}
