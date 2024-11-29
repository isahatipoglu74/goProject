package main

import (
	"fmt"

	"math/rand"
)

func main() {

	for x := 0; x <= 100; x++ {

		a := rand.Intn(10)
		b := rand.Intn(10)

		fmt.Printf("iteration %v x, y values %v %v ", x, a, b)
		switch {
		case a < 4 && b < 4:
			fmt.Println("a ve b 4 den küçüktür.")
		case a < 6 && b < 6:
			fmt.Println("a ve b 6 den büyüktür.")
		case a >= 4 && b <= 6:
			fmt.Println("a büuük 4 den b 6 dan kçük")
		case b != 5:
			fmt.Println("b 5 değildir")
		default:
			fmt.Println("Ocak dışı ")
		}
	}

}
