package main

import "fmt"

func main() {
	// key: value
	choi := map[string]string{"name": "choi", "age": "12"}
	fmt.Println(choi)
	for key, value := range choi {
		fmt.Println(key, value)
		fmt.Println(key)
		fmt.Println(value)
	}
}
