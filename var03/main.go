package main

import "fmt"

func main() {
	var a int16 = 3456   // 16비트 정수형 변수 선언
	var b int8 = int8(a) // 8비트 정수형 변수로 변환

	fmt.Println(a, b)
}
