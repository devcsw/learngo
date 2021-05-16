package main

import (
	"fmt"
	"strings"
)

//리턴 값 먼저 설정 하는법 (naked return)
func lenAndUpper(name string) (lenght int, uppercase string) {
	//함수가 끝난후 실행 defer
	defer fmt.Println("I'm done")
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

//여러 argument 받아들이는법
//배열로 받아짐
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println(lenAndUpper("dd"))

	totalLenght, upperName := lenAndUpper("choi")
	//defer로 끝났을때마다 실행
	fmt.Println(totalLenght, upperName)
	//여러값 리턴 무시하기 언더바(_)
	totalLenght2, _ := lenAndUpper("park")
	fmt.Println(totalLenght2)

	repeatMe("1", "2", "3", "4", "5")
}
