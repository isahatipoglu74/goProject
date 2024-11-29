package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {

		a := rand.Intn(5)

		switch a {
		case 0:
			fmt.Printf("iteration: %v \t a values %v \n", i, a)
		case 1:
			fmt.Printf("iteration: %v \t a values %v \n", i, a)
		case 2:
			fmt.Printf("iteration: %v \t a values %v \n", i, a)
		case 3:
			fmt.Printf("iteration: %v \t a values %v \n", i, a)
		case 4:
			fmt.Printf("iteration: %v \t a values %v \n", i, a)
		}
	}

}
