package main

import (
	"fmt"

	"github.com/isahatipoglu74/puppy"
)

func main() {

	fmt.Println("Helllo")

	s1 := puppy.Bark()
	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

}
