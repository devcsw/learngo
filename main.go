package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

//여러값 반환
func lenAndUpper(name string) (int, string, int) {
	return len(name), strings.ToUpper(name), 1
}

//여러 argument 받아들이는법
//배열로 받아짐
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(multiply(2, 2))
	fmt.Println(lenAndUpper("dd"))

	totalLenght, upperName, num := lenAndUpper("choi")
	fmt.Println(totalLenght, upperName, num)
	//여러값 리턴 무시하기 언더바(_)
	totalLenght2, upperName2, _ := lenAndUpper("park")
	fmt.Println(totalLenght2, upperName2)

	repeatMe("1", "2", "3", "4", "5")
}
