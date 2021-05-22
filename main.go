package main

import "fmt"

func main() {
	// names := [5]string{"choi", "lyn", "jung"}
	// names[3] = "1"
	// names[4] = "1"
	// names[5] = "1"

	//slice
	names := []string{"choi", "lyn", "jung"}

	names = append(names, "e")

	fmt.Println(names)
}
