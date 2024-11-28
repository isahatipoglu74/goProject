package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("Fonkisyon main func dan önce inititalize edildi.")
}

func main() {

	x := rand.Intn(250)
	fmt.Printf("The value of x is %v\t", x)

	if x <= 100 {
		fmt.Println("less than 100")
	} else if x >= 101 && x <= 200 {
		fmt.Println("101 ile 200 arasındadır")
	} else {
		fmt.Println("200 den büyüktür.")
	}

	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))

	switch {
	case x <= 100:
		fmt.Println("x = 100 den küçük")
	case x > 100 && x <= 200:
		fmt.Println("x 100 ile 200 arasında")
	case x > 201 && x <= 250:
		fmt.Println("x 200 ile 250 arasındadır")
	default:
		fmt.Println("x 250 den büyüktür.")
	}

	a := rand.Intn(10)
	b := rand.Intn(10)

	fmt.Printf("type %T , values %v \n", a, a)
	fmt.Printf("type %T , values %v \n", b, b)

	if a < 4 && b < 4 {
		fmt.Println("a ve b 4 den küçüktür.")
	} else if a < 6 && b < 6 {
		fmt.Println("a ve b 6 den büyüktür.")
	} else if a >= 4 && b <= 6 {
		fmt.Println("a büuük 4 den b 6 dan kçük")
	} else if b != 5 {
		fmt.Println("b 5 değildir")
	} else {
		fmt.Println("Ocak dışı ")
	}

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
