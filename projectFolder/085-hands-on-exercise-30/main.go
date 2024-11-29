package main

import (
	"fmt"
	"math/rand"
)

func main() {

	xi := []int{25, 23, 63, 44, 65, 98}

	for i, v := range xi {
		fmt.Printf("index %v  value %v \n ", i, v)
	}
	fmt.Println("------------------------------")
	m := map[string]int{
		"isa":   32,
		"feyza": 28,
	}

	for k, v := range m {
		fmt.Printf("key : %v , \t value : %v\n", k, v)
	}

	fmt.Println("------------------------------")

	age1 := m["isa"]
	fmt.Println("Yaş ", age1)
	if v, ok := m["isa"]; ok {
		fmt.Println("isa var", v)
	}

	age2 := m["veli"] //yukarıdaki map de veli diye bir key yok
	fmt.Println("Yaş ", age2)
	if v, ok := m["veli"]; !ok {
		fmt.Println("veli yok", v)
	}

	fmt.Println("------------------------------")

	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("iteration %v \t x değeri %v tür.\n", i, x)
		}
	}

}
