package main

import "fmt"

func main() {

	x := 83 / 40
	y := 83 % 40
	fmt.Println(x, y)

	for i := 0; i < 100; i++ {

		if i%2 == 0 {
			fmt.Printf("Bu çift sayıdır \t %v \n", i)
		}
	}
}
