package main

import "fmt"

var g int = 10

func main() {
	var m int = 10
	{
		var n int = 20
		fmt.Println(m, n, g)
	}
	// fmt.Println(m, n) // 에러 발생
}
