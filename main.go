package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// struct 오브젝트 같은데 map보단 유연함
	//구조체?
	favFood := []string{"cola", "pizza"}
	choi := person{"choi", 32, favFood}
	fmt.Println(choi)
	choi = person{
		name:    "choi",
		age:     31,
		favFood: favFood,
	}
	fmt.Println(choi)
}
