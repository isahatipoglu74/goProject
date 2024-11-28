package main

import "fmt"

var x = 12

const y = 13

func main() {
	var a = 2
	const b = 3
	fmt.Printf("x is value : %v , and x is type of : %T \n", x, x)
	fmt.Printf("y is value : %v , and y is type of : %T \n", y, y)

	fmt.Printf("a is value : %v , and a is type of : %T \n", a, a)
	fmt.Printf("b is value : %v , and b is type of : %T \n", b, b)
}
