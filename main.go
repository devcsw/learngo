package main

import "fmt"

func main() {
	a := 2
	b := &a
	a = 10

	fmt.Println(&a, b)
	fmt.Println(*b)
	fmt.Println(a)
	//b= 1
	*b = 20
	fmt.Println(&a, b)
	fmt.Println(*b)
	fmt.Println(a)
}
